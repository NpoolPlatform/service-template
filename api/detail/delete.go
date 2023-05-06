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

func (s *Server) DeleteDetail(ctx context.Context, in *npool.DeleteDetailRequest) (*npool.DeleteDetailResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteDetail", "ID", in.GetID(), "Error", err)
		return &npool.DeleteDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteDetail", "ID", in.GetID(), "Error", err)
		return &npool.DeleteDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

*/
