package detail

import (
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	entdetail "github.com/NpoolPlatform/service-template/pkg/db/ent/detail"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	SampleCol *string
}

func CreateSet(c *ent.DetailCreate, req *Req) *ent.DetailCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	EntID     *cruder.Cond
	SampleCol *cruder.Cond
}

func SetQueryConds(q *ent.DetailQuery, conds *Conds) (*ent.DetailQuery, error) {
	if conds.EntID != nil {
		switch conds.EntID.Op {
		case cruder.EQ:
			id, ok := conds.EntID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid ent id")
			}
			q.Where(entdetail.EntID(id))
		default:
			return nil, fmt.Errorf("invalid sample field")
		}
	}
	if conds.SampleCol != nil {
		switch conds.SampleCol.Op {
		case cruder.LIKE:
			sampleCol, ok := conds.SampleCol.Val.(string)
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
