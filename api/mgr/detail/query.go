//nolint:nolintlint,dupl
package detail

import (
	"context"

	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	converter "github.com/NpoolPlatform/service-template/pkg/mgr/detail/converter"
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

func (s *Server) GetDetail(ctx context.Context, in *npool.GetDetailRequest) (*npool.GetDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetDetail", "ID", in.GetID(), "Error", err)
		return &npool.GetDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "detail", "crud", "Row")

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
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetailOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "detail", "crud", "RowOnly")

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
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetDetail", "Conds", in.GetConds(), "Error", err)
		return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "detail", "crud", "Rows")

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
