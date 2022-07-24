package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/servicetmpl"

	"github.com/NpoolPlatform/service-template/api/detail"
	"github.com/NpoolPlatform/service-template/api/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	servicetmpl.UnimplementedServiceTemplateServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterServiceTemplateServer(server, &Server{})
	general.Register(server)
	detail.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := servicetmpl.RegisterServiceTemplateHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := general.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
