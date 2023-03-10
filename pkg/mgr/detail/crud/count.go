package detail

import (
	"context"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"

	tracer "github.com/NpoolPlatform/service-template/pkg/mgr/detail/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
)

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
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

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}
