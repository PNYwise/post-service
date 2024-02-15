package repository

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/PNYwise/post-service/internal/domain"
)

type kafkaPostRepository struct {
	producer sarama.SyncProducer
	extConf  *domain.ExtConf
}

func NewKafkaPostRepository(producer sarama.SyncProducer, extConf *domain.ExtConf) domain.KafkaPostRepository {
	return &kafkaPostRepository{
		producer: producer,
		extConf:  extConf,
	}
}

// PublishMessage implements domain.KafkaPostRepository.
func (k *kafkaPostRepository) PublishMessage(post *domain.Post) error {
	jsonMessage, err := json.Marshal(post)
	if err != nil {
		fmt.Printf("error to mashal %v", err)
		return err
	}
	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: k.extConf.Kafka.Topic,
		Value: sarama.ByteEncoder(jsonMessage),
	})
	return err
}
