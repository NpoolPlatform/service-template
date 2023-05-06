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

func (s *Server) GetDetail(ctx context.Context, in *npool.GetDetailRequest) (*npool.GetDetailResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithID(ctx, &in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDetail",
			"In", in,
			"Error", err,
		)
		return &npool.GetDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetDetail(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDetail",
			"In", in,
			"Error", err,
		)
		return &npool.GetDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDetails(ctx context.Context, in *npool.GetDetailsRequest) (*npool.GetDetailsResponse, error) {
	handler, err := detail1.NewHandler(
		ctx,
		detail1.WithConds(ctx, in.GetConds()),
		detail1.WithOffset(ctx, in.GetOffset()),
		detail1.WithLimit(ctx, in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDetails",
			"In", in,
			"Error", err,
		)
		return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetDetails(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDetails",
			"In", in,
			"Error", err,
		)
		return &npool.GetDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
