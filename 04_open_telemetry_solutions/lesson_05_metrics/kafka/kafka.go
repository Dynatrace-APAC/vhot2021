package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
)

// BrokerURL has no documentation
// LESSION 03: INSTRUMENTATION LIBRARIES
const BrokerURL = "kafka.mushroom.home:9092"

// TopicName has no documentation
const TopicName = "SomeTopic"

// Send pushes a message with the contents of the integer value into kafka
func Send(result int) {
	if producer, err := newSyncProducer([]string{BrokerURL}); err == nil {
		defer producer.Close()
		send(producer, result)
	}
}

// newSyncProducer has no documentation
func newSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	producer = otelsarama.WrapSyncProducer(config, producer)
	return producer, nil
}

func send(producer sarama.SyncProducer, msgID int) {
	msg := sarama.ProducerMessage{
		Topic:   TopicName,
		Key:     sarama.StringEncoder(fmt.Sprintf("%d", msgID)),
		Value:   sarama.StringEncoder("MessageValue"),
		Headers: []sarama.RecordHeader{{Key: []byte("testKeyHeader"), Value: []byte("testKeyValue")}},
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		return
	}

	fmt.Printf("Sent message, id: %d, partition: %d, offset: %d \n", msgID, partition, offset)

}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/

// PreserveImport is never getting called.
// In order to keep this session as simple as possible it's purpose is to preserve imports on top of this file
func PreserveImport() {
	otelsarama.WrapSyncProducer(nil, nil)
}
