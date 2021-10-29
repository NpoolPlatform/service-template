package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*npool.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &npool.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}
