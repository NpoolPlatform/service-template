package main

import (
	"time"

	"github.com/NpoolPlatform/service-template/api"
	db "github.com/NpoolPlatform/service-template/pkg/db"
	msgcli "github.com/NpoolPlatform/service-template/pkg/message/client"
	msglistener "github.com/NpoolPlatform/service-template/pkg/message/listener"
	msg "github.com/NpoolPlatform/service-template/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/service-template/pkg/message/server"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

const MsgInterval = 3 * time.Second

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()

		if err := msgsrv.Init(); err != nil {
			return err
		}
		if err := msgcli.Init(); err != nil {
			return err
		}

		go msglistener.Listen()
		go msgSender()

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apimgrcli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	apimgrcli.Register(mux)

	return nil
}

func msgSender() {
	id := 0
	for {
		logger.Sugar().Infof("send example")
		err := msgsrv.PublishExample(&msg.Example{
			ID:      id,
			Example: "hello world",
		})
		if err != nil {
			logger.Sugar().Errorf("fail to send example: %v", err)
			return
		}
		id++
		time.Sleep(MsgInterval)
	}
}
