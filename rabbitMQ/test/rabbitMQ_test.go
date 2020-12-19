package test

import (
	"encoding/json"
	"fmt"
	"github.com/spjiang/go-util/rabbitMQ"
	"github.com/streadway/amqp"
	"testing"
)

type cfg struct {
	UserName string
	Password string
	Host     string
	Port     int64
	Vhost    string
}

type deliverInout struct {
	routingKeys  []string // key值
	exchangeName string   // 交换机名称
	exchangeType string   // 交换机类型
}

var inoutConfig = &deliverInout{
	routingKeys:  []string{"inout.capture.new"},
	exchangeName: "parking.inout",
	exchangeType: "topic",
}

var mq *rabbitMQ.RMQ

// TestDeliver 消息投递
func TestDeliver(t *testing.T) {
	cfg := &cfg{
		UserName: "admin",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     15672,
		Vhost:    "/",
	}
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Vhost)
	mq = rabbitMQ.NewRMQ(addr)
	for _, routingKey := range inoutConfig.routingKeys {
		queueExchange := rabbitMQ.QueueExchange{
			RoutingKey:   routingKey,
			ExchangeName: inoutConfig.exchangeName,
			ExchangeType: inoutConfig.exchangeType,
		}
		paper := mq.InitPaper(queueExchange)
		paper.Pub([]byte("message"))
	}
}

type TestConsumer struct{}
type TestMsg struct {
	Name string
}

func (d *TestConsumer) Consumer(msg amqp.Delivery) error {
	defer func() {
		_ = msg.Ack(false)
	}()
	revData := &TestMsg{}
	err := json.Unmarshal(msg.Body, revData)
	if err != nil {
		return err
	}
	fmt.Println(revData.Name)
	return nil
}


// TestSubscribe 消费订阅
func TestSubscribe(t *testing.T) {
	cfg := &cfg{
		UserName: "admin",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     15672,
		Vhost:    "/",
	}
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Vhost)
	mq = rabbitMQ.NewRMQ(addr)

	queueExchange := rabbitMQ.QueueExchange{
		QueueName:    "TestQueueName",
		RoutingKey:   "new",
		ExchangeName: "TestName",
		ExchangeType: "topic",
	}
	paper := mq.InitPaper(queueExchange)
	consumer := &TestConsumer{}
	paper.Receive(consumer)
}
