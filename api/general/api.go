package general

import (
	"github.com/NpoolPlatform/message/npool/servicetmpl/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
