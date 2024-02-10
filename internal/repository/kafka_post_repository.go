package repository

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/PNYwise/post-service/internal/domain"
)

type kafkaPostRepository struct {
	producer sarama.SyncProducer
}

func NewKafkaPostRepository(producer sarama.SyncProducer) domain.KafkaPostRepository {
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
	if _, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "post",
		Value: sarama.ByteEncoder(jsonMessage),
	}); err != nil {
		fmt.Printf("Error executing kafka send message: %v", err)
		return err
	}
	return err
}
