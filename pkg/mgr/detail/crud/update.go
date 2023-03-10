package detail

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/NpoolPlatform/service-template/pkg/db/ent"
)

func UpdateSet(info *ent.Detail, in *npool.DetailReq) *ent.DetailUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.DetailReq) (*ent.Detail, error) {
	return nil, nil
}
