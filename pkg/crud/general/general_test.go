package general

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/service-template/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/general"
	testinit "github.com/NpoolPlatform/service-template/pkg/testinit"
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

var entity = ent.General{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	UserID:     uuid.New(),
	CoinTypeID: uuid.New(),
	Incoming:   decimal.NewFromInt(0),
	Locked:     decimal.NewFromInt(0),
	Outcoming:  decimal.NewFromInt(0),
	Spendable:  decimal.NewFromInt(0),
}

var (
	id         = entity.ID.String()
	appID      = entity.AppID.String()
	userID     = entity.UserID.String()
	coinTypeID = entity.CoinTypeID.String()
	incoming   = entity.Incoming.String()
	locked     = entity.Locked.String()
	outcoming  = entity.Outcoming.String()
	spendable  = entity.Spendable.String()

	req = npool.GeneralReq{
		ID:         &id,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Incoming:   &incoming,
		Locked:     &locked,
		Outcoming:  &outcoming,
		Spendable:  &spendable,
	}
)

var info *ent.General

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.General{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			Incoming:   decimal.NewFromInt(0),
			Locked:     decimal.NewFromInt(0),
			Outcoming:  decimal.NewFromInt(0),
			Spendable:  decimal.NewFromInt(0),
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			Incoming:   decimal.NewFromInt(0),
			Locked:     decimal.NewFromInt(0),
			Outcoming:  decimal.NewFromInt(0),
			Spendable:  decimal.NewFromInt(0),
		},
	}

	reqs := []*npool.GeneralReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_incoming := _entity.Incoming.String()
		_locked := _entity.Locked.String()
		_outcoming := _entity.Outcoming.String()
		_spendable := _entity.Spendable.String()

		reqs = append(reqs, &npool.GeneralReq{
			ID:         &_id,
			AppID:      &_appID,
			UserID:     &_userID,
			CoinTypeID: &_coinTypeID,
			Incoming:   &_incoming,
			Locked:     &_locked,
			Outcoming:  &_outcoming,
			Spendable:  &_spendable,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func add(t *testing.T) {
	incoming = "30"
	locked = "10"
	outcoming = "10"
	spendable = "10"

	req.Incoming = &incoming
	req.Locked = &locked
	req.Outcoming = &outcoming
	req.Spendable = &spendable

	entity.Incoming, _ = decimal.NewFromString(incoming)
	entity.Locked, _ = decimal.NewFromString(locked)
	entity.Outcoming, _ = decimal.NewFromString(outcoming)
	entity.Spendable, _ = decimal.NewFromString(spendable)

	info, err := AddFields(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0].String(), entity.String())
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestGeneral(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("add", add)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
