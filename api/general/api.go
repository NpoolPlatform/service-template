package general

import (
	"github.com/NpoolPlatform/message/npool/servicetmpl/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedTemplateGeneralServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterTemplateGeneralServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
