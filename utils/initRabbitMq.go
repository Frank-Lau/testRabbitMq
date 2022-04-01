package utils

//import (
//	"fmt"
//	"github.com/streadway/amqp"
//	"log"
//)
//
//const (
//	PublishExchange string = "dc_bank_fil_wallet_exchange"        //消息投递的交换机
//	BackupExchange  string = "dc_bank_fil_wallet_backup_exchange" //备份交换机
//	DeadExchange    string = "dc_bank_fil_wallet_dead_exchange"   //死信交换机
//	PublishQueue    string = "dc_bank_fil_wallet_queue"           //消息投递队列
//	BackupQueue     string = "dc_bank_fil_wallet_backup_queue"    //备份交换机队列
//	DeadQueue       string = "dc_bank_fil_wallet_dead_queue"      //死信队列
//	PublishKey      string = "dc_bank_fil_wallet_publish_key"     //消息投递rouking key
//)
//
////func PublishEvent(db *gorm.DB, amqpUri string, e *Event) error {
//func CreateRabbitMq() error {
//	conn := GetRabbitMqConn()
//	ch, err := conn.Channel()
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	defer ch.Close()
//	//声明备份交换器
//	if err = ch.ExchangeDeclare(
//		BackupExchange,
//		amqp.ExchangeDirect,
//		true,
//		false,
//		false,
//		false,
//		nil,
//	); err != nil {
//		log.Println("ch.ExchangeDeclare(备份交换机) err: ", err)
//		return err
//	}
//	//声明死信交换机
//	if err = ch.ExchangeDeclare(
//		DeadExchange,
//		amqp.ExchangeDirect,
//		true,
//		false,
//		false,
//		false,
//		nil,
//	); err != nil {
//		log.Println("ch.ExchangeDeclare(死信交换机) err: ", err)
//		return err
//	}
//	//添加备份交换器参数
//	argsExchange := make(map[string]interface{})
//	argsExchange["alternate-exchange"] = BackupExchange
//	//声明交换器并绑定备份交换器
//	if err = ch.ExchangeDeclare(
//		PublishExchange,
//		amqp.ExchangeDirect,
//		true,
//		false,
//		false,
//		false,
//		argsExchange, //指定备份交换器
//	); err != nil {
//		fmt.Println("ch.ExchangeDeclare(交换机) err: ", err)
//		return err
//	}
//	//var q amqp.Queue
//	args := make(map[string]interface{})
//	//设置队列的过期时间
//	//args["x-message-ttl"] = 10000
//	//设置死信交换器
//	args["x-dead-letter-exchange"] = DeadExchange
//	//设置死信交换器Key
//	args["x-dead-letter-routing-key"] = PublishKey
//	//创建实际使用的队列
//	exchangeQ, err := ch.QueueDeclare(
//		PublishQueue,
//		true,
//		false,
//		false,
//		false,
//		args,
//	)
//	if err != nil {
//		fmt.Println("创建实际使用的队列时发生错误,错误原因:", err)
//		return err
//	}
//	//为备份交换机声明队列
//	backupExchangeQ, err := ch.QueueDeclare(
//		BackupQueue,
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		fmt.Println("备份交换机队列声明的发生错误,错误原因:", err)
//		return err
//	}
//	//为死信交换机声明队列
//	deadxchangeQ, err := ch.QueueDeclare(
//		DeadQueue,
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	//绑定交换器/队列和key
//	if err = ch.QueueBind(exchangeQ.Name, PublishKey, PublishExchange, false, nil); err != nil {
//		fmt.Println("exchange  QueueBind err: ", err)
//		return err
//	}
//	//绑定备份交换机和队列
//	if err = ch.QueueBind(backupExchangeQ.Name, PublishKey, BackupExchange, false, nil); err != nil {
//		log.Println("backup QueueBind err: ", err)
//		return err
//	}
//	//绑定死信交换机和死信队列
//	if err = ch.QueueBind(deadxchangeQ.Name, PublishKey, DeadExchange, false, nil); err != nil {
//		log.Println("dead QueueBind err: ", err)
//		return err
//	}
//	return nil
//}
