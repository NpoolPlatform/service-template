package handle

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
)

func (s *Server) Echo(ctx context.Context, in *npool.StringMessage) (*npool.StringMessage, error) {
	return in, nil
}
