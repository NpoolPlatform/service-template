package detail

import (
	"github.com/NpoolPlatform/message/npool/servicetmpl/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
