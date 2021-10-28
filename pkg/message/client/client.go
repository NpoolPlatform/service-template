package client

import (
	"encoding/json"

	"golang.org/x/xerrors"

	constant "github.com/NpoolPlatform/go-service-app-template/pkg/message/const"
	msg "github.com/NpoolPlatform/go-service-app-template/pkg/message/message"
	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"

	"github.com/streadway/amqp"
)

type client struct {
	*msgcli.Client
	consumers map[string]<-chan amqp.Delivery
}

var myClients = map[string]*client{}

func Init() error {
	_myClient, err := msgcli.New(constant.ServiceName)
	if err != nil {
		return err
	}

	err = _myClient.DeclareQueue(msg.QueueExample)
	if err != nil {
		return err
	}

	sampleClient := &client{
		Client:    _myClient,
		consumers: map[string]<-chan amqp.Delivery{},
	}
	examples, err := _myClient.Consume(msg.QueueExample)
	if err != nil {
		return xerrors.Errorf("fail to construct example consume: %v", err)
	}
	sampleClient.consumers[msg.QueueExample] = examples

	myClients[constant.ServiceName] = sampleClient

	return nil
}

func ConsumeExample(h func(*msg.Example) error) error {
	examples, ok := myClients[constant.ServiceName].consumers[msg.QueueExample]
	if !ok {
		return xerrors.Errorf("consumer is not constructed")
	}

	for d := range examples {
		example := msg.Example{}
		err := json.Unmarshal(d.Body, &example)
		if err != nil {
			return xerrors.Errorf("parse message example error: %v", err)
		}

		if h != nil {
			err = h(&example)
			if err != nil {
				return err
			}
		}
	}

	return xerrors.Errorf("WE SHOULD NOT BE HERE")
}
