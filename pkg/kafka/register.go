package kafka

import (
	"errors"
	"fmt"
	"github.com/Hackaton-UDEVS/first/internal/config"
)

func Register(h *KafkaHandler) error {
	cfg := config.Load()
	brokers := []string{fmt.Sprintf("%s:%d", cfg.KAFKAHOST, cfg.KAFKAPORT)}
	kcm := NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "app-create", "app-create-id", h.NewAppCreate()); err != nil {
		if err == ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'books-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	return nil
}
