package detail

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	detail1 "github.com/NpoolPlatform/service-template/pkg/mw/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateDetail(ctx context.Context, in *npool.UpdateDetailRequest) (*npool.UpdateDetailResponse, error) {
	req := in.GetInfo()
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithEntID(ctx, req.EntID),
		detail1.WithSampleCol(ctx, req.SampleCol),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDetail",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateDetail(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDetail",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDetailResponse{
		Info: info,
	}, nil
}
