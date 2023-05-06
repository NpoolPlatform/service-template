package detail

/*
import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"

	"github.com/google/uuid"
)

func (s *Server) GetDetail(ctx context.Context, in *npool.GetDetailRequest) (*npool.GetDetailResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetDetail", "ID", in.GetID(), "Error", err)
		return &npool.GetDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetDetail", "ID", in.GetID(), "Error", err)
		return &npool.GetDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func ValidateConds(in *npool.Conds) error {
	return nil
}

func (s *Server) GetDetailOnly(ctx context.Context, in *npool.GetDetailOnlyRequest) (*npool.GetDetailOnlyResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDetails(ctx context.Context, in *npool.GetDetailsRequest) (*npool.GetDetailsResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

*/
