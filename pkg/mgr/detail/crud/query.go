package detail

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	tracer "github.com/NpoolPlatform/service-template/pkg/mgr/detail/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/detail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func Row(ctx context.Context, id string) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Detail.Query().Where(detail.ID(uuid.MustParse(id))).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.DetailQuery, error) { //nolint
	stm := cli.Detail.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOType != nil {
		switch conds.GetIOType().GetOp() {
		case cruder.EQ:
			stm.Where(detail.IoType(npool.IOType(conds.GetIOType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOSubType != nil {
		switch conds.GetIOSubType().GetOp() {
		case cruder.EQ:
			stm.Where(detail.IoType(npool.IOSubType(conds.GetIOSubType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Amount != nil {
		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(detail.AmountLT(amount))
		case cruder.GT:
			stm.Where(detail.AmountGT(amount))
		case cruder.EQ:
			stm.Where(detail.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.FromCoinTypeID != nil {
		switch conds.GetFromCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.FromCoinTypeID(uuid.MustParse(conds.GetFromCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(conds.GetCoinUSDCurrency().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetCoinUSDCurrency().GetOp() {
		case cruder.LT:
			stm.Where(detail.CoinUsdCurrencyLT(currency))
		case cruder.GT:
			stm.Where(detail.CoinUsdCurrencyGT(currency))
		case cruder.EQ:
			stm.Where(detail.CoinUsdCurrencyEQ(currency))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOExtra != nil {
		switch conds.GetIOExtra().GetOp() {
		case cruder.LIKE:
			stm.Where(detail.IoExtraContains(conds.GetIOExtra().GetValue()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.FromOldID != nil {
		switch conds.GetFromOldID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.FromOldID(uuid.MustParse(conds.GetFromOldID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Detail, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Detail{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(detail.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
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

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
