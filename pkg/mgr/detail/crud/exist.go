package detail

import (
	"context"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	tracer "github.com/NpoolPlatform/service-template/pkg/mgr/detail/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/detail"

	"github.com/google/uuid"
)

func Exist(ctx context.Context, id string) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Detail.Query().Where(detail.ID(uuid.MustParse(id))).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
