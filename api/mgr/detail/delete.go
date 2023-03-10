//nolint:nolintlint,dupl
package detail

import (
	"context"

	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	converter "github.com/NpoolPlatform/service-template/pkg/mgr/detail/converter"
	crud "github.com/NpoolPlatform/service-template/pkg/mgr/detail/crud"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/google/uuid"
)

func (s *Server) DeleteDetail(ctx context.Context, in *npool.DeleteDetailRequest) (*npool.DeleteDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteDetail", "ID", in.GetID(), "Error", err)
		return &npool.DeleteDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceID(span, in.GetID())
	span = commontracer.TraceInvoker(span, "detail", "crud", "Delete")

	info, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteDetail", "ID", in.GetID(), "Error", err)
		return &npool.DeleteDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
