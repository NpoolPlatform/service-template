//nolint:nolintlint,dupl
package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/service-template/pkg/crud/template"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	constant "github.com/NpoolPlatform/service-template/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/template"

	"github.com/google/uuid"
)

func checkTemplateInfo(info *npool.TemplateReq) error {
	if info.Name == nil {
		logger.Sugar().Error("Name is empty")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	if info.Age == nil {
		logger.Sugar().Error("Age is empty")
		return status.Error(codes.InvalidArgument, "Age is empty")
	}

	return nil
}

func templateRowToObject(row *ent.Template) *npool.Template {
	return &npool.Template{
		ID:   row.ID.String(),
		Name: row.Name,
		Age:  row.Age,
	}
}

func (s *TemplateServer) CreateTemplate(ctx context.Context, in *npool.CreateTemplateRequest) (*npool.CreateTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TemplateSpanAttributes(span, in.GetInfo())
	err = checkTemplateInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateTemplateResponse{}, err
	}

	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create template: %v", err)
		return &npool.CreateTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTemplateResponse{
		Info: templateRowToObject(info),
	}, nil
}

func (s *TemplateServer) CreateTemplates(ctx context.Context, in *npool.CreateTemplatesRequest) (*npool.CreateTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTemplates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTemplatesResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}

	dup := make(map[string]struct{})
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("ID.%v", key), info.GetID()),
			attribute.String(fmt.Sprintf("Name.%v", key), info.GetName()),
			attribute.String(fmt.Sprintf("Age.%v", key), info.GetName()),
		)
		err := checkTemplateInfo(info)
		if err != nil {
			return &npool.CreateTemplatesResponse{}, err
		}

		if _, ok := dup[info.GetName()]; ok {
			return &npool.CreateTemplatesResponse{},
				status.Errorf(codes.AlreadyExists,
					"Name: %v duplicate create",
					info.GetName(),
				)
		}

		dup[info.GetName()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create templates: %v", err)
		return &npool.CreateTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.Template, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, templateRowToObject(val))
	}

	return &npool.CreateTemplatesResponse{
		Infos: infos,
	}, nil
}

func (s *TemplateServer) UpdateTemplate(ctx context.Context, in *npool.UpdateTemplateRequest) (*npool.UpdateTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TemplateSpanAttributes(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("template id is invalid")
		return &npool.UpdateTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update template: %v", err)
		return &npool.UpdateTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTemplateResponse{
		Info: templateRowToObject(info),
	}, nil
}

func (s *TemplateServer) GetTemplate(ctx context.Context, in *npool.GetTemplateRequest) (*npool.GetTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get template: %v", err)
		return &npool.GetTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTemplateResponse{
		Info: templateRowToObject(info),
	}, nil
}

func (s *TemplateServer) GetTemplateOnly(ctx context.Context, in *npool.GetTemplateOnlyRequest) (*npool.GetTemplateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTemplateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TemplateCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get templates: %v", err)
		return &npool.GetTemplateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTemplateOnlyResponse{
		Info: templateRowToObject(info),
	}, nil
}

func (s *TemplateServer) GetTemplates(ctx context.Context, in *npool.GetTemplatesRequest) (*npool.GetTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTemplates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TemplateCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)

	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get templates: %v", err)
		return &npool.GetTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.Template, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, templateRowToObject(val))
	}

	return &npool.GetTemplatesResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *TemplateServer) ExistTemplate(ctx context.Context, in *npool.ExistTemplateRequest) (*npool.ExistTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check template: %v", err)
		return &npool.ExistTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTemplateResponse{
		Info: exist,
	}, nil
}

func (s *TemplateServer) ExistTemplateConds(ctx context.Context,
	in *npool.ExistTemplateCondsRequest) (*npool.ExistTemplateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTemplateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TemplateCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check template: %v", err)
		return &npool.ExistTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTemplateCondsResponse{
		Info: exist,
	}, nil
}

func (s *TemplateServer) CountTemplates(ctx context.Context, in *npool.CountTemplatesRequest) (*npool.CountTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTemplates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.TemplateCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count templates: %v", err)
		return &npool.CountTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTemplatesResponse{
		Info: total,
	}, nil
}

func (s *TemplateServer) DeleteTemplate(ctx context.Context, in *npool.DeleteTemplateRequest) (*npool.DeleteTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete template: %v", err)
		return &npool.DeleteTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTemplateResponse{
		Info: templateRowToObject(info),
	}, nil
}
