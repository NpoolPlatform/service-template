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

func (s *Server) CreateDetail(ctx context.Context, in *npool.CreateDetailRequest) (*npool.CreateDetailResponse, error) {
	req := in.GetInfo()
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithEntID(ctx, req.EntID),
		detail1.WithSampleCol(ctx, req.SampleCol),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDetail",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateDetail(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDetail",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateDetails(ctx context.Context, in *npool.CreateDetailsRequest) (*npool.CreateDetailsResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithReqs(ctx, in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDetails",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateDetails(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDetails",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailsResponse{
		Infos: infos,
	}, nil
}
