package main

import (
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/go-service-app-template/api"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		go func() {
			err := grpc2.RunGRPC(rpcRegister)
			if err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()
		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return api.RegisterGateway(mux, endpoint, opts)
}
