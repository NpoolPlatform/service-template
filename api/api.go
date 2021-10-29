package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedServiceExampleServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterServiceExampleServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterServiceExampleHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
