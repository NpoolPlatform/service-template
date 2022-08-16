package detail

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
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/detail"
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

var entity = ent.Detail{
	ID:              uuid.New(),
	AppID:           uuid.New(),
	UserID:          uuid.New(),
	CoinTypeID:      uuid.New(),
	IoType:          npool.IOType_Incoming.String(),
	IoSubType:       npool.IOSubType_Payment.String(),
	Amount:          decimal.RequireFromString("9999999999999999999.999999999999999999"),
	FromCoinTypeID:  uuid.New(),
	CoinUsdCurrency: decimal.RequireFromString("1.00045000000123012"),
	IoExtra:         uuid.New().String(),
	FromOldID:       uuid.New(),
}

var (
	id              = entity.ID.String()
	appID           = entity.AppID.String()
	userID          = entity.UserID.String()
	coinTypeID      = entity.CoinTypeID.String()
	ioType          = npool.IOType(npool.IOType_value[entity.IoType])
	ioSubType       = npool.IOSubType(npool.IOSubType_value[entity.IoSubType])
	amount          = entity.Amount.String()
	fromCoinTypeID  = entity.FromCoinTypeID.String()
	coinUSDCurrency = entity.CoinUsdCurrency.String()
	ioExtra         = entity.IoExtra
	fromOldID       = entity.FromOldID.String()

	req = npool.DetailReq{
		ID:              &id,
		AppID:           &appID,
		UserID:          &userID,
		CoinTypeID:      &coinTypeID,
		IOType:          &ioType,
		IOSubType:       &ioSubType,
		Amount:          &amount,
		FromCoinTypeID:  &fromCoinTypeID,
		CoinUSDCurrency: &coinUSDCurrency,
		IOExtra:         &ioExtra,
		FromOldID:       &fromOldID,
	}
)

var info *ent.Detail

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
	entities := []*ent.Detail{
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			UserID:          uuid.New(),
			CoinTypeID:      uuid.New(),
			IoType:          npool.IOType_Incoming.String(),
			IoSubType:       npool.IOSubType_Payment.String(),
			Amount:          decimal.RequireFromString("10.00896"),
			FromCoinTypeID:  uuid.New(),
			CoinUsdCurrency: decimal.RequireFromString("1.8902"),
			IoExtra:         uuid.New().String(),
			FromOldID:       uuid.New(),
		},
		{
			ID:              uuid.New(),
			AppID:           uuid.New(),
			UserID:          uuid.New(),
			CoinTypeID:      uuid.New(),
			IoType:          npool.IOType_Incoming.String(),
			IoSubType:       npool.IOSubType_Payment.String(),
			Amount:          decimal.RequireFromString("11.11111"),
			FromCoinTypeID:  uuid.New(),
			CoinUsdCurrency: decimal.RequireFromString("1.123"),
			IoExtra:         uuid.New().String(),
			FromOldID:       uuid.New(),
		},
	}

	reqs := []*npool.DetailReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_ioType := npool.IOType(npool.IOType_value[_entity.IoType])
		_ioSubType := npool.IOSubType(npool.IOSubType_value[_entity.IoSubType])
		_amount := _entity.Amount.String()
		_fromCoinTypeID := entity.FromCoinTypeID.String()
		_coinUSDCurrency := _entity.CoinUsdCurrency.String()
		_ioExtra := _entity.IoExtra
		_fromOldID := _entity.FromOldID.String()

		reqs = append(reqs, &npool.DetailReq{
			ID:              &_id,
			AppID:           &_appID,
			UserID:          &_userID,
			CoinTypeID:      &_coinTypeID,
			IOType:          &_ioType,
			IOSubType:       &_ioSubType,
			Amount:          &_amount,
			FromCoinTypeID:  &_fromCoinTypeID,
			CoinUSDCurrency: &_coinUSDCurrency,
			IOExtra:         &_ioExtra,
			FromOldID:       &_fromOldID,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
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
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
		}
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

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
