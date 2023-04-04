package subscriber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/leoantony72/twitter-backend/timeline/internal/model"
	"github.com/leoantony72/twitter-backend/timeline/internal/ports"
	"github.com/leoantony72/twitter-backend/timeline/pkg"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var Msgs chan *amqp.Delivery

// var done chan string

func ConsumeTweets(ch *amqp.Channel, repo ports.TimelineRepository) {
	QNAME := pkg.GetEnv("QNAME")
	CONSUMER := pkg.GetEnv("CONSUMER")
	Msgs, err := ch.Consume(QNAME, CONSUMER, true, false, false, false, nil)
	failOnError(err, "failed to setup consumer")

	for {
		msg := <-Msgs
		go Process(&msg, repo)
	}

}

func Process(msg *amqp.Delivery, repo ports.TimelineRepository) {
	tweet, err := deserialize(msg.Body)
	fmt.Println(tweet)
	if err != nil {
		fmt.Println(err)
	}
	user := tweet.Username
	followers := repo.GetFollowers(user)

	length := len(followers)
	if length < 10 {
		//push method
		for _, username := range followers {
			fmt.Println(username)
			err := repo.AddTimelineEntry(tweet, username)
			if err != nil {
				continue
			}
		}
	}
	//pull method

}

func deserialize(b []byte) (model.Tweets, error) {
	msg := model.Tweets{}
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
