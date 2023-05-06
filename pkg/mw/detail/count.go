package detail

import (
	"context"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"

	detailcrud "github.com/NpoolPlatform/service-template/pkg/crud/detail"
)

func (h *Handler) CountDetails(ctx context.Context) (uint32, error) {
	count := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := detailcrud.SetQueryConds(cli.Detail.Query(), h.Conds)
		if err != nil {
			return err
		}
		_count, err := stm.Count(_ctx)
		if err != nil {
			return err
		}
		count = uint32(_count)
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}
