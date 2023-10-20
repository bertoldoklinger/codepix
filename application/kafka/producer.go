package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}
	producer, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return producer
}


func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: []byte(msg),
	}
	err:= producer.Produce(message, deliveryChan)

	if err!= nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	// loopamos e ficamos escutando tudo oq chega no meu channel, rodando pra sempre
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery Failed", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}	
}