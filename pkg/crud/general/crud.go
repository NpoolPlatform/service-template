package general

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"
	tracer "github.com/NpoolPlatform/service-template/pkg/tracer/general"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/general"
	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/general"

	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.GeneralReq) (*ent.General, error) {
	var info *ent.General
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
		c := cli.General.Create()

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

		c.SetIncoming(decimal.NewFromInt(0))
		c.SetLocked(decimal.NewFromInt(0))
		c.SetOutcoming(decimal.NewFromInt(0))
		c.SetSpendable(decimal.NewFromInt(0))

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.GeneralReq) ([]*ent.General, error) {
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

	rows := []*ent.General{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.GeneralCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.General.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.CoinTypeID != nil {
				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetCoinTypeID()))
			}
			bulk[i].SetIncoming(decimal.NewFromInt(0))
			bulk[i].SetLocked(decimal.NewFromInt(0))
			bulk[i].SetOutcoming(decimal.NewFromInt(0))
			bulk[i].SetSpendable(decimal.NewFromInt(0))
		}
		rows, err = tx.General.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func AddFields(ctx context.Context, in *npool.GeneralReq) (*ent.General, error) { //nolint
	var info *ent.General
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

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.General.Query().Where(general.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query general: %v", err)
		}

		incoming := decimal.NewFromInt(0)
		if in.Incoming != nil {
			amount, err := decimal.NewFromString(in.GetIncoming())
			if err != nil {
				return err
			}
			incoming = incoming.Add(amount)
		}
		locked := decimal.NewFromInt(0)
		if in.Locked != nil {
			amount, err := decimal.NewFromString(in.GetLocked())
			if err != nil {
				return err
			}
			locked = locked.Add(amount)
		}
		outcoming := decimal.NewFromInt(0)
		if in.Outcoming != nil {
			amount, err := decimal.NewFromString(in.GetOutcoming())
			if err != nil {
				return err
			}
			outcoming = outcoming.Add(amount)
		}
		spendable := decimal.NewFromInt(0)
		if in.Spendable != nil {
			amount, err := decimal.NewFromString(in.GetSpendable())
			if err != nil {
				return err
			}
			spendable = spendable.Add(amount)
		}

		if incoming.Add(info.Incoming).
			Cmp(
				locked.Add(info.Locked).
					Add(outcoming).
					Add(info.Outcoming).
					Add(spendable).
					Add(info.Spendable),
			) < 0 {
			return fmt.Errorf("outcoming (%v + %v) + locked (%v + %v) + spendable (%v + %v) > incoming (%v + %v)",
				outcoming, info.Outcoming, locked, info.Locked, spendable, info.Spendable, incoming, info.Incoming)
		}

		if locked.Add(info.Locked).Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("locked + locked < 0")
		}

		if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("incoming < 0")
		}

		if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("outcoming < 0")
		}

		if spendable.Add(info.Spendable).Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("spendable + spendable < 0")
		}

		stm := info.Update()

		if in.Incoming != nil {
			stm = stm.AddIncoming(incoming)
		}
		if in.Outcoming != nil {
			stm = stm.AddOutcoming(outcoming)
		}
		if in.Locked != nil {
			stm = stm.AddLocked(locked)
		}
		if in.Spendable != nil {
			stm = stm.AddSpendable(spendable)
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update general: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update general: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.General.Query().Where(general.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.GeneralQuery, error) { //nolint
	stm := cli.General.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(general.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(general.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(general.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(general.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.Incoming != nil {
		incoming, err := decimal.NewFromString(conds.GetIncoming().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetIncoming().GetOp() {
		case cruder.LT:
			stm.Where(general.IncomingLT(incoming))
		case cruder.GT:
			stm.Where(general.IncomingGT(incoming))
		case cruder.EQ:
			stm.Where(general.IncomingEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.Locked != nil {
		locked, err := decimal.NewFromString(conds.GetLocked().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetLocked().GetOp() {
		case cruder.LT:
			stm.Where(general.LockedLT(locked))
		case cruder.GT:
			stm.Where(general.LockedGT(locked))
		case cruder.EQ:
			stm.Where(general.LockedEQ(locked))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.Outcoming != nil {
		outcoming, err := decimal.NewFromString(conds.GetOutcoming().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetOutcoming().GetOp() {
		case cruder.LT:
			stm.Where(general.OutcomingLT(outcoming))
		case cruder.GT:
			stm.Where(general.OutcomingGT(outcoming))
		case cruder.EQ:
			stm.Where(general.OutcomingEQ(outcoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.Spendable != nil {
		spendable, err := decimal.NewFromString(conds.GetSpendable().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSpendable().GetOp() {
		case cruder.LT:
			stm.Where(general.SpendableLT(spendable))
		case cruder.GT:
			stm.Where(general.SpendableGT(spendable))
		case cruder.EQ:
			stm.Where(general.SpendableEQ(spendable))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.General, int, error) {
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

	rows := []*ent.General{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(general.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.General, error) {
	var info *ent.General
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
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

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
		stm, err := setQueryConds(conds, cli)
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

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
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

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.General.Query().Where(general.ID(id)).Exist(_ctx)
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
		stm, err := setQueryConds(conds, cli)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.General.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
