package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// criar um canal de evento do kafka
	// se a mensagem foi entregue ou não
	deliveryChan := make(chan kafka.Event)

	producer := NewKafkaProducer()

	// usando go routines
	Publish("Mensagem teste", "teste", producer, nil, deliveryChan)

	go DeliveryReport(deliveryChan)

	// forma síncrona de resolver
	// e := <- deliveryChan
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar mensagem")
	// } else {
	// 	fmt.Println("Mensagem enviada:", msg.TopicPartition)
	// }

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
func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	// monta a struct do kafka para enviar a mensagem
	message := &kafka.Message{
		Value: []byte(msg), // pode enviar qualquer tipo de informação
		TopicPartition: kafka.TopicPartition{ Topic: &topic, Partition: kafka.PartitionAny },
		Key: key,
	}

	err := producer.Produce(message, deliveryChan) // o resultado dessa mensagem é publicada no canal
	if err != nil {
		return err
	}

	return nil
}


// função com looping infinito para pegar o retorno do canal
func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("Erro ao enviar mensagem")
				} else {
					fmt.Println("Mensagem enviada:", ev.TopicPartition)
					// Pode ser criado algo
				}
		}
	}
}