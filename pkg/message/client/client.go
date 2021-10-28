package client

import (
	"encoding/json"

	"golang.org/x/xerrors"

	msg "github.com/NpoolPlatform/go-service-app-template/pkg/message/message"
	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	"github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/common" //nolint
)

var myClient *msgcli.Client

func Init() error {
	_myClient, err := msgcli.New(rabbitmq.MyServiceNameToVHost())
	if err != nil {
		return err
	}

	myClient = _myClient

	return nil
}

func ConsumeExample() ([]*msg.Example, error) {
	msgs, err := myClient.Consume(msg.QueueExample)
	if err != nil {
		return nil, xerrors.Errorf("consume example error: %v", err)
	}

	examples := []*msg.Example{}
	for d := range msgs {
		example := &msg.Example{}
		err := json.Unmarshal(d.Body, &example)
		if err != nil {
			return nil, xerrors.Errorf("parse message example error: %v", err)
		}
		examples = append(examples, example)
	}

	return examples, nil
}
