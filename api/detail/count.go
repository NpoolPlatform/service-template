package detail

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	detail1 "github.com/NpoolPlatform/service-template/pkg/mw/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CountDetails(ctx context.Context, in *npool.CountDetailsRequest) (*npool.CountDetailsResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithConds(ctx, in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountDetails",
			"In", in,
			"Error", err,
		)
		return &npool.CountDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := handler.CountDetails(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountDetails",
			"In", in,
			"Error", err,
		)
		return &npool.CountDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountDetailsResponse{
		Info: count,
	}, nil
}
