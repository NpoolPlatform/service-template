package server

import (
	msg "github.com/NpoolPlatform/go-service-app-template/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
)

func Init() error {
	err := msgsrv.Init()
	if err != nil {
		return err
	}

	err = msgsrv.DeclareQueue(msg.QueueExample)
	if err != nil {
		return err
	}

	return nil
}

var Deinit = msgsrv.Deinit                 //nolint
var PublishToQueue = msgsrv.PublishToQueue //nolint
