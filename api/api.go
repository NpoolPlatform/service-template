package api

import (
	"context"

	servicetmpl "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1"

	"github.com/NpoolPlatform/service-template/api/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	servicetmpl.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterMiddlewareServer(server, &Server{})
	detail.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := servicetmpl.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
