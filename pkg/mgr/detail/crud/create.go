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

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateSet(c *ent.DetailCreate, in *npool.DetailReq) *ent.DetailCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.IOType != nil {
		c.SetIoType(in.GetIOType().String())
	}
	if in.IOSubType != nil {
		c.SetIoSubType(in.GetIOSubType().String())
	}
	if in.Amount != nil {
		c.SetAmount(decimal.RequireFromString(in.GetAmount()))
	}
	if in.FromCoinTypeID != nil {
		c.SetFromCoinTypeID(uuid.MustParse(in.GetFromCoinTypeID()))
	}
	if in.CoinUSDCurrency != nil {
		c.SetCoinUsdCurrency(decimal.RequireFromString(in.GetCoinUSDCurrency()))
	}
	if in.IOExtra != nil {
		c.SetIoExtra(in.GetIOExtra())
	}
	if in.FromOldID != nil {
		c.SetFromOldID(uuid.MustParse(in.GetFromOldID()))
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	return c
}

func Create(ctx context.Context, in *npool.DetailReq) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Detail.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.DetailReq) ([]*ent.Detail, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Detail{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.DetailCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.Detail.Create(), info)
		}
		rows, err = cli.Detail.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
