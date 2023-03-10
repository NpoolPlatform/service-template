//nolint:nolintlint,dupl
package detail

import (
	"context"

	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	crud "github.com/NpoolPlatform/service-template/pkg/mgr/detail/crud"
	tracer "github.com/NpoolPlatform/service-template/pkg/mgr/detail/tracer"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"
)

func (s *Server) CountDetails(ctx context.Context, in *npool.CountDetailsRequest) (*npool.CountDetailsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("CountDetails", "Conds", in.GetConds(), "Error", err)
		return &npool.CountDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "detail", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountDetails", "Conds", in.GetConds(), "Error", err)
		return &npool.CountDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountDetailsResponse{
		Info: total,
	}, nil
}
