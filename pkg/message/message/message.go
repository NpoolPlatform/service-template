package message

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
)

const (
	QueueExample = "example"
)

type Example struct {
	ID      int    `json:"id"`
	Example string `json:"example"`
}

func InitQueues() error {
	err := msgsrv.DeclareQueue(QueueExample)
	if err != nil {
		return err
	}
	return nil
}
