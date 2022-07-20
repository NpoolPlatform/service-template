package template

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/service-template/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	val "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/template"
	testinit "github.com/NpoolPlatform/service-template/pkg/test-init"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entTemplate = ent.Template{
	ID:   uuid.New(),
	Name: uuid.New().String(),
	Age:  10,
}

var (
	id           = entTemplate.ID.String()
	templateInfo = npool.TemplateReq{
		ID:   &id,
		Name: &entTemplate.Name,
		Age:  &entTemplate.Age,
	}
)

var info *ent.Template

func rowToObject(row *ent.Template) *ent.Template {
	return &ent.Template{
		ID:   row.ID,
		Name: row.Name,
		Age:  row.Age,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &templateInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entTemplate.ID = info.ID
			entTemplate.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entTemplate)
	}
}

func createBulk(t *testing.T) {
	entTemplate := []ent.Template{
		{
			ID:   uuid.New(),
			Name: uuid.New().String(),
			Age:  10,
		},
		{
			ID:   uuid.New(),
			Name: uuid.New().String(),
			Age:  10,
		},
	}

	templates := []*npool.TemplateReq{}
	for key := range entTemplate {
		id := entTemplate[key].ID.String()
		templates = append(templates, &npool.TemplateReq{
			ID:   &id,
			Name: &entTemplate[key].Name,
			Age:  &entTemplate[key].Age,
		})
	}
	infos, err := CreateBulk(context.Background(), templates)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &templateInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTemplate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteA)
	t.Run("count", count)
}
