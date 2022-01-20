package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/go-service-app-template"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedServiceExampleServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterServiceExampleServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterServiceExampleHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
