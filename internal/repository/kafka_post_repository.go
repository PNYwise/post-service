package repository

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/PNYwise/post-service/internal/domain"
)

type kafkaPostRepository struct {
	producer sarama.AsyncProducer
}

func NewKafkaPostRepository(producer sarama.AsyncProducer) domain.KafkaPostRepository {
	return &kafkaPostRepository{
		producer: producer,
	}
}

// PublishMessage implements domain.KafkaPostRepository.
func (k *kafkaPostRepository) PublishMessage(post *domain.Post) error {
	jsonMessage, err := json.Marshal(post)
	if err != nil {
		fmt.Printf("error mashaling post: %v", err)
		return err
	}
	// Construct the ProducerMessage
	msg := &sarama.ProducerMessage{
		Topic: "post",
		Value: sarama.ByteEncoder(jsonMessage),
	}
	// Send the message and handle any errors
	select {
	case k.producer.Input() <- msg:
		// Message sent successfully
		return nil
	case err := <-k.producer.Errors():
		fmt.Printf("Error occurred while sending the message: %v", err)
		return err
	}
}
