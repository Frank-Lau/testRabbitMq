package main

import (
	"fmt"
	"math/rand"
	"testRabbitMq/producer"
	"time"
)

func main() {
	//go consumer.RunTaskQueueConsumer()
	startTime := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		uidInt := i
		packetId := i + 1
		emailPhone := 18812341000 + i
		money := rand.Intn(1000)
		p := make([]interface{}, 0)
		p = append(p, uidInt)
		p = append(p, packetId)
		p = append(p, emailPhone)
		p = append(p, money)
		//producer.PublishEvent(db, configs.GetAmqpUrlConf(), utils.NewEvent(utils.EventRedPacket,
		producer.PublishEvent("amqp://guest:guest@192.168.12.251:5672/", producer.NewEvent("red_packet", "", "", p, time.Now().Unix()))
	}
	endTime := time.Now().UnixNano()
	fmt.Println("用时:", endTime-startTime)
}
