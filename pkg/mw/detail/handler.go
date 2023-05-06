package detail

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	constant "github.com/NpoolPlatform/service-template/pkg/const"
	detailcrud "github.com/NpoolPlatform/service-template/pkg/crud/detail"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	AutoID    *uint32
	ID        *uuid.UUID
	SampleCol *string
	Reqs      []*detailcrud.Req
	Conds     *detailcrud.Conds
	Offset    int32
	Limit     int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(ctx context.Context, id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithSampleCol(ctx context.Context, sampleCol *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SampleCol = sampleCol
		return nil
	}
}

func WithReqs(ctx context.Context, reqs []*npool.DetailReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Reqs = []*detailcrud.Req{}
		for _, req := range reqs {
			_req := &detailcrud.Req{
				SampleCol: req.SampleCol,
			}
			if req.ID != nil {
				id, err := uuid.Parse(req.GetID())
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			h.Reqs = append(h.Reqs, _req)
		}
		return nil
	}
}

func WithConds(ctx context.Context, conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &detailcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.AutoID != nil {
			h.Conds.AutoID = &cruder.Cond{
				Op:  conds.GetAutoID().GetOp(),
				Val: conds.GetAutoID().GetValue(),
			}
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.SampleCol != nil {
			h.Conds.SampleCol = &cruder.Cond{
				Op:  conds.GetSampleCol().GetOp(),
				Val: conds.GetSampleCol().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(ctx context.Context, offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(ctx context.Context, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
