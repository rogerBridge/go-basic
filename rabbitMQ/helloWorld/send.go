package main

import (
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to rabbitmq server")
	defer conn.Close() // 最先defer的函数最后执行
	// 建立信道
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 建立queue和往队列中丢信息, durable
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	FailOnError(err, "failed to declare a queue")
	// 往queue里面推送信息
	for i := 0; i < 10000; i++ {
		// 在信道上建立queue and publish
		body := "第" + strconv.Itoa(i) + "条消息@" + time.Now().Format("2006-01-02 15:04:05.999")
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
		FailOnError(err, "Failed to publish a message")
	}
}

// FailOnError return err+msg
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
