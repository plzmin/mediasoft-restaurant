package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"restaurant/internal/model"
	"restaurant/internal/repository/orderrepository"
	"restaurant/pkg/logger"
)

type Consumer struct {
	log             *logger.Logger
	consumer        sarama.Consumer
	orderRepository orderrepository.OrderRepository
}

func New(brokers []string, log *logger.Logger, orderRepository orderrepository.OrderRepository) (*Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, cfg)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		log:             log,
		consumer:        consumer,
		orderRepository: orderRepository,
	}, err
}

func (c *Consumer) Consume(topic string) {

	c.log.Info("Consuming partition of topic %s", topic)
	partitionList, err := c.consumer.Partitions(topic)
	if err != nil {
		c.log.Fatal("Failed to start consumer for partition %v", err)
	}
	for _, partition := range partitionList {
		pc, _ := c.consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				order := new(model.Order)
				if err = json.Unmarshal(message.Value, order); err != nil {
					c.log.Error("failed Unmarshal msg %v", err.Error())
				}
				if err = c.orderRepository.Create(order); err != nil {
					c.log.Error("failed save msg %v", err.Error())
				}
			}
		}(pc)
	}
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
