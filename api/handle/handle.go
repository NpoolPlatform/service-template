package handle

import (
	"github.com/NpoolPlatform/go-service-app-template/message/npool"
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
