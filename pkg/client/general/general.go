//nolint:dupl
package general

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/general"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.TemplateGeneralClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get general connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewTemplateGeneralClient(conn)

	return handler(_ctx, cli)
}

func CreateGeneral(ctx context.Context, in *npool.GeneralReq) (*npool.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.CreateGeneral(ctx, &npool.CreateGeneralRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create general: %v", err)
	}
	return info.(*npool.General), nil
}

func CreateGenerals(ctx context.Context, in []*npool.GeneralReq) ([]*npool.General, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.CreateGenerals(ctx, &npool.CreateGeneralsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create generals: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create generals: %v", err)
	}
	return infos.([]*npool.General), nil
}

func AddGeneral(ctx context.Context, in *npool.GeneralReq) (*npool.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.AddGeneral(ctx, &npool.AddGeneralRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update general: %v", err)
	}
	return info.(*npool.General), nil
}

func GetGeneral(ctx context.Context, id string) (*npool.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.GetGeneral(ctx, &npool.GetGeneralRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get general: %v", err)
	}
	return info.(*npool.General), nil
}

func GetGeneralOnly(ctx context.Context, conds *npool.Conds) (*npool.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.GetGeneralOnly(ctx, &npool.GetGeneralOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get general: %v", err)
	}
	return info.(*npool.General), nil
}

func GetGenerals(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.General, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.GetGenerals(ctx, &npool.GetGeneralsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get generals: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get generals: %v", err)
	}
	return infos.([]*npool.General), total, nil
}

func ExistGeneral(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.ExistGeneral(ctx, &npool.ExistGeneralRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get general: %v", err)
	}
	return infos.(bool), nil
}

func ExistGeneralConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.ExistGeneralConds(ctx, &npool.ExistGeneralCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get general: %v", err)
	}
	return infos.(bool), nil
}

func CountGenerals(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.CountGenerals(ctx, &npool.CountGeneralsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count general: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteGeneral(ctx context.Context, id string) (*npool.General, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.TemplateGeneralClient) (cruder.Any, error) {
		resp, err := cli.DeleteGeneral(ctx, &npool.DeleteGeneralRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete general: %v", err)
	}
	return infos.(*npool.General), nil
}
