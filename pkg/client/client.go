//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/template"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
)

var timeout = 10 * time.Second

func doTemplate(ctx context.Context,
	fn func(_ctx context.Context,
		cli npool.ServiceTemplateTemplateClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get template connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewServiceTemplateTemplateClient(conn)

	return fn(_ctx, cli)
}

func CreateTemplate(ctx context.Context, in *npool.TemplateReq) (*npool.Template, error) {
	info, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.CreateTemplate(ctx, &npool.CreateTemplateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create template: %v", err)
	}
	return info.(*npool.Template), nil
}

func CreateTemplates(ctx context.Context, in []*npool.TemplateReq) ([]*npool.Template, error) {
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.CreateTemplates(ctx, &npool.CreateTemplatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create templates: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create templates: %v", err)
	}
	return infos.([]*npool.Template), nil
}

func UpdateTemplate(ctx context.Context, in *npool.TemplateReq) (*npool.Template, error) {
	info, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.UpdateTemplate(ctx, &npool.UpdateTemplateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update template: %v", err)
	}
	return info.(*npool.Template), nil
}

func GetTemplate(ctx context.Context, id string) (*npool.Template, error) {
	info, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.GetTemplate(ctx, &npool.GetTemplateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get template: %v", err)
	}
	return info.(*npool.Template), nil
}

func GetTemplateOnly(ctx context.Context, conds *npool.Conds) (*npool.Template, error) {
	info, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.GetTemplateOnly(ctx, &npool.GetTemplateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get template: %v", err)
	}
	return info.(*npool.Template), nil
}

func GetTemplates(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Template, uint32, error) {
	var total uint32
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.GetTemplates(ctx, &npool.GetTemplatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get templates: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get templates: %v", err)
	}
	return infos.([]*npool.Template), total, nil
}

func ExistTemplate(ctx context.Context, id string) (bool, error) {
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.ExistTemplate(ctx, &npool.ExistTemplateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get template: %v", err)
	}
	return infos.(bool), nil
}

func ExistTemplateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.ExistTemplateConds(ctx, &npool.ExistTemplateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get template: %v", err)
	}
	return infos.(bool), nil
}

func CountTemplates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.CountTemplates(ctx, &npool.CountTemplatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count template: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteTemplate(ctx context.Context, id string) (*npool.Template, error) {
	infos, err := doTemplate(ctx, func(_ctx context.Context, cli npool.ServiceTemplateTemplateClient) (cruder.Any, error) {
		resp, err := cli.DeleteTemplate(ctx, &npool.DeleteTemplateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete template: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete template: %v", err)
	}
	return infos.(*npool.Template), nil
}
