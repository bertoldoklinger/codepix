package cmd

import (
	"os"

	"github.com/bertoldoklinger/codepix/application/kafka"
	"github.com/bertoldoklinger/codepix/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)


var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		deliveryChan := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()
		
		// roda de forma assincrona escutando com go routines/conccorrencia
		go kafka.DeliveryReport(deliveryChan)
		

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},

}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}