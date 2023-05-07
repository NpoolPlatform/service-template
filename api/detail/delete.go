package detail

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	detail1 "github.com/NpoolPlatform/service-template/pkg/mw/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteDetail(ctx context.Context, in *npool.DeleteDetailRequest) (*npool.DeleteDetailResponse, error) {
	req := in.GetInfo()
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithEntID(ctx, req.EntID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDetail",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteDetail(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDetail",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDetailResponse{
		Info: info,
	}, nil
}
