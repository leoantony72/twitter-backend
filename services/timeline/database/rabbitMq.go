package database

import (
	"log"

	"github.com/leoantony72/twitter-backend/timeline/pkg"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func StartMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	failOnError(err, "failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	//defer ch.Close()

	ExchName := pkg.GetEnv("EXCHNAME")
	err = ch.ExchangeDeclare(ExchName, "direct", true, false, false, false, nil)
	failOnError(err, "failed to declare exchange")

	qName := pkg.GetEnv("QNAME")
	q, err := ch.QueueDeclare(qName, true, false, false, false, nil)
	failOnError(err, "failed to declared queue")
	err = ch.QueueBind(q.Name, "", ExchName, false, nil)
	failOnError(err, "failed to bind queue")

	// ch.Publish(ExchName,"",false,false,amqp.Publishing{
	// 	ContentType: "text/plain",
	// 	Body: []byte("sample"),
	// })

	return ch

}
