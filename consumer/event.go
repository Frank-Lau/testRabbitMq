package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"testRabbitMq/producer"
	"testRabbitMq/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}
func RunTaskQueueConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.12.251:5672/")
	failOnError(err, "Failed to dial url")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"red_packet", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	//消费者内存中只允许有1000条为ack的消息
	err = ch.Qos(
		1000,  // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	for d := range msgs {
		event := new(producer.Event)
		err = json.Unmarshal(d.Body, event)
		fmt.Println("body===>>>", string(d.Body))
		failOnError(err, "Decode json error")
		switch event.Name {

		case "rd_packet":
			logrus.Printf("Received a message: %s", d.Body)
			uid, _ := utils.CheckNumberAndConvertToInt(event.Params[0])
			packetId := event.Params[1].(string)
			emailPhone := event.Params[2].(string)
			drawAmount := event.Params[3].(string)
			isNew, _ := utils.CheckNumberAndConvertToInt(event.Params[4])
			//rt := DrawPacketEvent(uid, packetId, emailPhone, drawAmount, isNew)
			//if rt.Err != nil {
			//	logrus.Error(rt.Err)
			//}
			fmt.Printf("uid:%d,packetId:%d,emailPhone:%s,drawAmount:%s,isNew:%d", uid, packetId, emailPhone, drawAmount, isNew)
		}
		d.Ack(false) //确认接收
		//d.Nack(false, false) //拒绝消费
	}
	<-forever
}
