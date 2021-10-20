package echo

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedServiceExampleServer
}

func (s *Server) Echo(ctx context.Context, in *npool.StringMessage) (*npool.StringMessage, error) {
	return in, nil
}
