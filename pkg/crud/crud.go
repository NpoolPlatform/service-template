package detail

import (
	"context"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	entdetail "github.com/NpoolPlatform/service-template/pkg/db/ent/detail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uuid.UUID
	SampleCol *string
}

func CreateSet(c *ent.DetailCreate, req *Req) *ent.DetailCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.SampleCol != nil {
		c.SetSampleCol(*req.SampleCol)
	}
	return c
}

func UpdateSet(u *ent.DetailUpdateOne, req *Req) *ent.DetailUpdateOne {
	if req.SampleCol != nil {
		u.SetSampleCol(*req.SampleCol)
	}
	return u
}

type Conds struct {
	AutoID    *cruder.Cond
	ID        *cruder.Cond
	SampleCol *cruder.Cond
}

func SetQueryConds(q *ent.DetailQuery, conds *Conds) (*ent.DetailQuery, error) {
	if conds.AutoID != nil {
		switch conds.AutoID.Op {
		case cruder.EQ:
			q.Where(entdetail.AutoID(conds.AutoID.Val))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	if conds.ID != nil {
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entdetail.ID(conds.ID.Val))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	if conds.SampleCol != nil {
		switch conds.SampleCol.Op {
		case cruder.LIKE:
			q.Where(entdetail.SampleCol(conds.SampleCol.Val))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
}
