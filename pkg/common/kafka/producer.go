package kafkacommon

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/edge-api/pkg/models"
	log "github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}

var singleInstance *kafka.Producer

// GetProducerInstance returns a kafka producer instance
func GetProducerInstance() *kafka.Producer {
	log.Debug("Getting the producer instance")
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		cfg := config.Get()
		if cfg.KafkaBrokers != nil {
			log.WithFields(log.Fields{"broker": cfg.KafkaBrokers[0].Hostname,
				"port": *cfg.KafkaBrokers[0].Port}).Debug("Creating a new producer")
			// FIXME: the ConfigMap should honor omitempty to avoid nil reference panics for unused features
			/*			p, err := kafka.NewProducer(&kafka.ConfigMap{
						"bootstrap.servers": fmt.Sprintf("%s:%d", cfg.KafkaBrokers[0].Hostname, *cfg.KafkaBrokers[0].Port),
						"sasl.mechanisms":   brokers[0].Sasl.SaslMechanism,
						"security.protocol": brokers[0].Sasl.SecurityProtocol,
						"sasl.username":     brokers[0].Sasl.Username,
						"sasl.password":     brokers[0].Sasl.Password}) */
			p, err := kafka.NewProducer(&kafka.ConfigMap{
				"bootstrap.servers": fmt.Sprintf("%s:%v", cfg.KafkaBrokers[0].Hostname, *cfg.KafkaBrokers[0].Port)})
			if err != nil {
				log.WithField("error", err).Error("Failed to create producer")
				return nil
			}
			singleInstance = p
		}
	}
	return singleInstance
}

// ProduceEvent is a helper for the kafka producer
func ProduceEvent(requestedTopic, recordKey string, event models.CRCCloudEvent) error {
	log.Debug("Producing an event")
	producer := GetProducerInstance()
	if producer == nil {
		log.Error("Failed to get the producer instance")
	}
	realTopic, err := GetTopic(requestedTopic)
	if err != nil {
		log.WithField("error", err).Error("Unable to lookup requested topic name")
	}

	// marshal the event into a string
	edgeEventMessage, err := json.Marshal(event)
	if err != nil {
		log.Error("Marshal CRCCloudEvent failed")
	}
	log.WithField("event", string(edgeEventMessage)).Debug("Debug CRCCloudEvent contents")

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &realTopic, Partition: kafka.PartitionAny},
		Key:            []byte(recordKey),
		Value:          edgeEventMessage,
	}, nil)
	if err != nil {
		log.WithField("error", err.Error()).Debug("Failed to produce the event")
		return err
	}

	return nil
}
