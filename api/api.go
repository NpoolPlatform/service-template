package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedServiceTemplateServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterServiceTemplateServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterServiceTemplateHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
