package producer

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"log"
)

const (
	EventRedPacket = "red_packet" //红包
)

type Event struct {
	Name     string        `json:"name"`
	Ip       string        `json:"ip"`
	Endpoint string        `json:"endpoint"`
	Params   []interface{} `json:"params"`
	Time     int64         `json:"time"`
}

func NewEvent(name string, ip string, endpoint string, params []interface{}, time int64) *Event {
	return &Event{
		Name:     name,
		Ip:       ip,
		Endpoint: endpoint,
		Params:   params,
		Time:     time,
	}
}

type EventLog struct {
	gorm.Model
	Name     string
	Ip       string
	Endpoint string
	Params   string
}

//func PublishEvent(db *gorm.DB, amqpUri string, e *Event) error {
func PublishEvent(amqpUri string, e *Event) error {
	conn, err := amqp.Dial(amqpUri)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	//go func() {
	//	eLog := new(EventLog)
	//	eLog.Name = e.Name
	//	eLog.Endpoint = e.Endpoint
	//	p, _ := json.Marshal(e.Params)
	//	eLog.Params = string(p)
	//
	//	db.Save(eLog)
	//}()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()
	//使用confirm模式,防止生产者丢消息
	err = ch.Confirm(false)
	if err != nil {
		fmt.Printf("error: %s \n", err.Error())
		return err
	}
	//声明备份交换器(这里的备份交换机没有设置绑定的队列,许手动或另外添加代码实现)
	if err = ch.ExchangeDeclare(
		"backup_exchange",
		amqp.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		fmt.Println("ch.ExchangeDeclare(备份交换机) err: ", err)
	}
	argsExchange := make(map[string]interface{})
	//添加备份交换器参数
	argsExchange["alternate-exchange"] = "backup_exchange"

	//声明交换器
	if err = ch.ExchangeDeclare(
		"long_direct",
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		argsExchange, //指定备份交换器
	); err != nil {
		fmt.Println("ch.ExchangeDeclare(交换机) err: ", err)
	}
	var q amqp.Queue
	args := make(map[string]interface{})
	//设置队列的过期时间
	args["x-message-ttl"] = 10000
	//设置死信交换器
	args["x-dead-letter-exchange"] = "exchange.dlx"
	//设置死信交换器Key
	args["x-dead-letter-routing-key"] = "dlxKey"
	switch e.Name {
	default:
		q, err = ch.QueueDeclare(
			"red_packet",
			true,
			false,
			false,
			false,
			nil,
		)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	//消费者confirm模式
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	defer confirmOne(confirms)
	//绑定交换器/队列和key
	if err = ch.QueueBind(q.Name, "red_packet", "long_direct", false, nil); err != nil {
		fmt.Println("ch.QueueBind err: ", err)
	}

	body, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ch.Publish(
		"long_direct", // exchange
		//q.Name, // routing key
		"red", // routing key 找不到匹配队列时将会被推送到备份交换机
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Printf("send message to rabbitmq %s\n", string(body))
	return nil
}
func confirmOne(confirms <-chan amqp.Confirmation) {
	if confirmed := <-confirms; confirmed.Ack {
		fmt.Printf("confirmed delivery with delivery tag: %d\n", confirmed.DeliveryTag)
	} else {
		fmt.Printf("confirmed delivery of delivery tag: %d\n", confirmed.DeliveryTag)
	}
}
