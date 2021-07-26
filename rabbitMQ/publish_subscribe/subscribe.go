package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 为了程序的健壮性, exchange声明了两次
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// bind exchange and queue
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // receive
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a receive")

	// 匿名函数 一个将string放入chan, 一个从chan里面拿string
	s := make(chan string, 10000)
	// 与main goroutine 并行的执行
	go func() {
		for d := range msgs {
			s <- string(d.Body)
			//log.Printf(" [x] %s", d.Body)
		}
	}()

	go func() {
		for i := range s{
			log.Println("",i)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	// forever长度为0, 没有输入, 不可能输出, 可以让main goroutine不退出
	forever := make(chan bool)
	<-forever
}
