package producer

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

func WriteBatteryTime(batteryLevel *device.BatteryLevel, deviceId string, uid string) error {
	manager := config.GetInstance()

	saramaConfig := manager.GetKafkaSarama()
	
	saramaConfig.Consumer.Return.Errors = true
	brokers := manager.GetBrokerList()
	
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = manager.GetMaxRetry()
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		logger.Error("Can't create SyncProducer", zap.Error(err))
		return err
	}

	defer func() {
		if err := producer.Close(); err != nil {
			logger.Error("Can't close", zap.Error(err))
		}
	}()

	marshaledData, err := json.Marshal(batteryLevel)
	if err != nil {
		logger.Error("Can't Marshal data", zap.Error(err))
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "battery_device__" + uid + "__" + deviceId,
		Value: sarama.StringEncoder(marshaledData),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)
	return nil
}
