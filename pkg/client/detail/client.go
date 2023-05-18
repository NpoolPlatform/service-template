//nolint:dupl
package detail

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"

	servicename "github.com/NpoolPlatform/service-template/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get detail connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateDetail(ctx context.Context, in *npool.DetailReq) (*npool.Detail, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateDetail(ctx, &npool.CreateDetailRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create detail: %v", err)
	}
	return info.(*npool.Detail), nil
}

func CreateDetails(ctx context.Context, in []*npool.DetailReq) ([]*npool.Detail, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateDetails(ctx, &npool.CreateDetailsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create details: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create details: %v", err)
	}
	return infos.([]*npool.Detail), nil
}

func GetDetail(ctx context.Context, id string) (*npool.Detail, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDetail(ctx, &npool.GetDetailRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get detail: %v", err)
	}
	return info.(*npool.Detail), nil
}

func GetDetailOnly(ctx context.Context, conds *npool.Conds) (*npool.Detail, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDetailOnly(ctx, &npool.GetDetailOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get detail: %v", err)
	}
	return info.(*npool.Detail), nil
}

func GetDetails(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Detail, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDetails(ctx, &npool.GetDetailsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get details: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get details: %v", err)
	}
	return infos.([]*npool.Detail), total, nil
}

func ExistDetail(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistDetail(ctx, &npool.ExistDetailRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get detail: %v", err)
	}
	return infos.(bool), nil
}

func ExistDetailConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistDetailConds(ctx, &npool.ExistDetailCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get detail: %v", err)
	}
	return infos.(bool), nil
}

func CountDetails(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CountDetails(ctx, &npool.CountDetailsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count detail: %v", err)
	}
	return infos.(uint32), nil
}
