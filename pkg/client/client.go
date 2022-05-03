package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.ServiceTemplateClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get service template connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewServiceTemplateClient(conn)

	return fn(_ctx, cli)
}

func GetServiceTemplateInfoOnly(ctx context.Context, conds cruder.FilterConds) (*npool.ServiceTemplateInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ServiceTemplateClient) (cruder.Any, error) {
		// DO RPC CALL HERE WITH conds PARAMETER
		return &npool.ServiceTemplateInfo{}, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get service template: %v", err)
	}
	return info.(*npool.ServiceTemplateInfo), nil
}
