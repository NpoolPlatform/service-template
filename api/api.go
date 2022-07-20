package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/servicetmpl/template"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type TemplateServer struct {
	template.UnimplementedServiceTemplateTemplateServer
}

func Register(server grpc.ServiceRegistrar) {
	template.RegisterServiceTemplateTemplateServer(server, &TemplateServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return template.RegisterServiceTemplateTemplateHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
