package main

import (
	"context"

	"github.com/NpoolPlatform/service-template/api"
	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/migrator"

	"github.com/NpoolPlatform/service-template/pkg/feeder"
	"github.com/NpoolPlatform/service-template/pkg/pubsub"
	"github.com/NpoolPlatform/service-template/pkg/watcher"

	action "github.com/NpoolPlatform/go-service-framework/pkg/action"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

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
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
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

func shutdown(ctx context.Context) {
	<-ctx.Done()
	logger.Sugar().Infow(
		"Watch",
		"State", "Done",
		"Error", ctx.Err(),
	)
	_ = pubsub.Shutdown(ctx) //nolint
}

func _watch(ctx context.Context, cancel context.CancelFunc, w func(ctx context.Context)) {
	defer func() {
		if err := recover(); err != nil {
			logger.Sugar().Errorw(
				"Watch",
				"State", "Panic",
				"Error", err,
			)
			cancel()
		}
	}()
	w(ctx)
}

func watch(ctx context.Context, cancel context.CancelFunc) error {
	go shutdown(ctx)
	go _watch(ctx, cancel, watcher.Watch)
	go _watch(ctx, cancel, feeder.Watch)
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
