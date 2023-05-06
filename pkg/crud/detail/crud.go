package detail

import (
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	entdetail "github.com/NpoolPlatform/service-template/pkg/db/ent/detail"
	"github.com/google/uuid"
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
			id, ok := conds.AutoID.Val.(uint32)
			if !ok {
				return nil, fmt.Errorf("invalid auto id")
			}
			q.Where(entdetail.AutoID(int(id)))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	if conds.ID != nil {
		switch conds.ID.Op {
		case cruder.EQ:
			id, ok := conds.ID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid id")
			}
			q.Where(entdetail.ID(id))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	if conds.SampleCol != nil {
		switch conds.SampleCol.Op {
		case cruder.LIKE:
			sampleCol, ok := conds.ID.Val.(string)
			if !ok {
				return nil, fmt.Errorf("invalid sample col")
			}
			q.Where(entdetail.SampleCol(sampleCol))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	q.Where(entdetail.DeletedAt(0))
	return q, nil
}
