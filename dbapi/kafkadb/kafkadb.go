package dbapi

import (
	"context"

	"github.com/EfimoffN/receivingLogs/config"
	"github.com/EfimoffN/receivingLogs/reciver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaAPI struct {
	kfk   *kafka.Producer
	topic string
}

func NewKafkaAPI(kfk *kafka.Producer, topic string) *KafkaAPI {
	return &KafkaAPI{
		kfk:   kfk,
		topic: topic,
	}
}

func ConnectKafka(cfg config.KafkaConfig) (*kafka.Producer, error) {
	host := "host1:" + cfg.LocalHost
	clientId := cfg.ClientID

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"client.id":         clientId, //socket.gethostname(),
		"acks":              "all"})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (api *KafkaAPI) SaveLog(ctx context.Context, sLog reciver.SendLog) error { // TODO: deal with kafka

	delivery_chan := make(chan kafka.Event, 10000)
	err := api.kfk.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &api.topic, Partition: kafka.PartitionAny},
		Value:          []byte("")},
		delivery_chan,
	)
	if err != nil {
		return err
	}

	return nil
}
