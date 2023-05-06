package detail

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"
)

func (h *Handler) GetDetail(ctx context.Context) (*npool.Detail, error) {
	if h.AutoID == nil {
		return nil, fmt.Errorf("invalid auto_id")
	}
	return nil, nil
}

func (h *Handler) GetDetails(ctx context.Context) ([]*npool.Detail, uint32, error) {
	return nil, 0, nil
}
