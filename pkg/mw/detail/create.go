package detail

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
	detailcrud "github.com/NpoolPlatform/service-template/pkg/crud/detail"
	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
)

func (h *Handler) CreateDetail(ctx context.Context) (*npool.Detail, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := detailcrud.CreateSet(
			cli.Detail.Create(),
			&detailcrud.Req{
				EntID:     h.EntID,
				SampleCol: h.SampleCol,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.EntID = &info.EntID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetDetail(ctx)
}

func (h *Handler) CreateDetails(ctx context.Context) ([]*npool.Detail, error) {
	return nil, nil
}
