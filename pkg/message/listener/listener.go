package listener

import (
	msgcli "github.com/NpoolPlatform/go-service-app-template/pkg/message/client"
	msg "github.com/NpoolPlatform/go-service-app-template/pkg/message/message"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func listenTemplateExample() {
	for {
		logger.Sugar().Infof("consume template example")
		err := msgcli.ConsumeExample(func(example *msg.Example) error {
			logger.Sugar().Infof("go example: %+w", example)
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
