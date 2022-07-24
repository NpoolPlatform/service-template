package listener

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	msgcli "github.com/NpoolPlatform/service-template/pkg/message/client"
	msg "github.com/NpoolPlatform/service-template/pkg/message/message"
)

func listenTemplateExample() {
	for {
		err := msgcli.ConsumeExample(func(example *msg.Example) error {
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}

func Listen() {
	go listenTemplateExample()
}
