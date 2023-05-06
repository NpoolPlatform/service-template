package detail

/*
import (
	"context"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"

	"github.com/google/uuid"
)

func (s *Server) ExistDetail(ctx context.Context, in *npool.ExistDetailRequest) (*npool.ExistDetailResponse, error) {

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("ExistDetail", "ID", in.GetID(), "Error", err)
		return &npool.ExistDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistDetail", "ID", in.GetID(), "Error", err)
		return &npool.ExistDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistDetailConds(ctx context.Context,
	in *npool.ExistDetailCondsRequest) (*npool.ExistDetailCondsResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("ExistDetailConds", "Conds", in.GetConds(), "Error", err)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistDetailConds", "Conds", in.GetConds(), "Error", err)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailCondsResponse{
		Info: exist,
	}, nil
}

*/
