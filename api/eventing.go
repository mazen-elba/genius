package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// Using Kafka for handling click events

func produceEvent(event *ClickEvent) {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "click_events",
		Value: sarama.StringEncoder(fmt.Sprintf("%v:%v", event.AdID, event.Timestamp)),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
}
