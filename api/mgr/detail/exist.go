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

	"github.com/google/uuid"
)

func (s *Server) ExistDetail(ctx context.Context, in *npool.ExistDetailRequest) (*npool.ExistDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("ExistDetail", "ID", in.GetID(), "Error", err)
		return &npool.ExistDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "detail", "crud", "Exist")

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
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDetailConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("ExistDetailConds", "Conds", in.GetConds(), "Error", err)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "detail", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistDetailConds", "Conds", in.GetConds(), "Error", err)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailCondsResponse{
		Info: exist,
	}, nil
}
