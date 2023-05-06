//nolint:dupl
package detail

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	detail1 "github.com/NpoolPlatform/service-template/pkg/mw/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistDetail(ctx context.Context, in *npool.ExistDetailRequest) (*npool.ExistDetailResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithID(ctx, &in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDetail",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistDetail(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDetail",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistDetailConds(ctx context.Context, in *npool.ExistDetailCondsRequest) (*npool.ExistDetailCondsResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithConds(ctx, in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDetailConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistDetailConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDetailConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailCondsResponse{
		Info: exist,
	}, nil
}
