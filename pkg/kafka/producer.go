package kafka

import (
	"crypto/sha512"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type Producer interface {
	SendMessage(topic string, key string, value interface{}) error
	Close() error
}

type saramaProducer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(cnf ProducerConfig) (Producer, error) {
	config := sarama.NewConfig()

	version, err := sarama.ParseKafkaVersion(cnf.Version)
	if err != nil {
		return nil, fmt.Errorf("error parsing Kafka version: %v", err)
	}
	config.Version = version

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Idempotent = true
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 200 * time.Millisecond
	config.Net.MaxOpenRequests = 1

	config.Producer.Compression = sarama.CompressionSnappy
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 30 * time.Second
	config.Net.WriteTimeout = 30 * time.Second
	config.Producer.Timeout = 10 * time.Second

	if cnf.SaslEnable {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = cnf.SaslUser
		config.Net.SASL.Password = cnf.SaslPassword
		config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: sha512.New}
		}
	}

	if cnf.TlsEnable {
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: false,
		}
	}

	producer, err := sarama.NewSyncProducer(cnf.Brokers, config)
	if err != nil {
		return nil, err
	}

	return &saramaProducer{syncProducer: producer}, nil
}

func (p *saramaProducer) SendMessage(topic string, key string, value interface{}) error {
	return p.sendMessageWithHeaders(topic, key, value, nil)
}

func (p *saramaProducer) sendMessageWithHeaders(topic string, key string, value interface{}, headers map[string]string) error {
	var payload []byte
	var err error

	switch v := value.(type) {
	case []byte:
		payload = v
	case string:
		payload = []byte(v)
	default:
		payload, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
		Headers: []sarama.RecordHeader{
			{Key: []byte("event-source"), Value: []byte("cmdb-system")},
		},
	}

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			msg.Headers = append(msg.Headers, sarama.RecordHeader{
				Key:   []byte(k),
				Value: []byte(v),
			})
		}
	}

	_, _, err = p.syncProducer.SendMessage(msg)
	return err
}

func (p *saramaProducer) Close() error {
	return p.syncProducer.Close()
}
