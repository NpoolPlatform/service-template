package main

import (
	"context"

	"github.com/NpoolPlatform/service-template/api"
	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/migrator"

	"github.com/NpoolPlatform/service-template/pkg/feeder"
	"github.com/NpoolPlatform/service-template/pkg/watcher"

	action "github.com/NpoolPlatform/go-service-framework/pkg/action"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		return action.Run(
			c.Context,
			func(ctx context.Context) error {
				return run(ctx)
			},
			rpcRegister,
			rpcGatewayRegister,
			func(ctx context.Context) error {
				return watch(ctx)
			},
		)
	},
}

func run(ctx context.Context) error {
	if err := migrator.Migrate(ctx); err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}
	return nil
}

func watch(ctx context.Context) error {
	go watcher.Watch(ctx)
	go feeder.Watch(ctx)
	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apicli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	_ = apicli.Register(mux)

	return nil
}
