package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// criar um canal de evento do kafka
	// se a mensagem foi entregue ou não
	deliveryChan := make(chan kafka.Event)


	producer := NewKafkaProducer()

	Publish("Mensagem teste", "teste", producer, nil)
	// a aplicação desliga antes de terminar o envio
	producer.Flush(5000)
}

// Apenas cria o producer
func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go_kafka-kafka-1:9092",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return producer
}



// Função de publicação de mensagens usando o producer
func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	// monta a struct do kafka para enviar a mensagem
	message := &kafka.Message{
		Value: []byte(msg), // pode enviar qualquer tipo de informação
		TopicPartition: kafka.TopicPartition{ Topic: &topic, Partition: kafka.PartitionAny },
		Key: key,
	}

	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}