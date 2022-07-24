package detail

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/detail"

	"github.com/google/uuid"
)

func validate(info *npool.DetailReq) error { //nolint
	if info.AppID == nil {
		logger.Sugar().Error("AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Error("UserID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Error("CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Error("CoinTypeID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.FromCoinTypeID != nil {
		if _, err := uuid.Parse(info.GetFromCoinTypeID()); err != nil {
			logger.Sugar().Error("FromCoinTypeID is invalid: %v", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("FromCoinTypeID is invalid: %v", err))
		}
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Error("Amount is invalid")
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Error("Amount is less than 0")
			return status.Error(codes.InvalidArgument, "Amount is less than 0")
		}
	}

	if info.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(info.GetCoinUSDCurrency())
		if err != nil {
			logger.Sugar().Error("CoinUSDCurrency is invalid")
			return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinUSDCurrency is invalid: %v", err))
		}
		if currency.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Error("CoinUSDCurrency is less than 0")
			return status.Error(codes.InvalidArgument, "CoinUSDCurrency is less than 0")
		}
	}

	if info.IOType == nil {
		logger.Sugar().Error("IOType is empty")
		return status.Error(codes.InvalidArgument, "IOType is empty")
	}

	if info.IOSubType == nil {
		logger.Sugar().Error("IOSubType is empty")
		return status.Error(codes.InvalidArgument, "IOSubType is empty")
	}

	switch info.GetIOType() {
	case npool.IOType_Incoming:
		switch info.GetIOSubType() {
		case npool.IOSubType_Payment:
		case npool.IOSubType_MiningBenefit:
		case npool.IOSubType_Commission:
		case npool.IOSubType_TechniqueFeeCommission:
		default:
			logger.Sugar().Error("Incoming IOSubType is invalid")
			return status.Error(codes.InvalidArgument, "Incoming IOSubType is invalid")
		}
	case npool.IOType_Outcoming:
		switch info.GetIOSubType() {
		case npool.IOSubType_Payment:
		case npool.IOSubType_Withdrawal:
		default:
			logger.Sugar().Error("Outcoming IOSubType is invalid")
			return status.Error(codes.InvalidArgument, "Outcoming IOSubType is invalid")
		}
	default:
		logger.Sugar().Error("IOSubType is invalid")
		return status.Error(codes.InvalidArgument, "IOSubType is invalid")
	}

	if info.FromOldID != nil {
		if _, err := uuid.Parse(info.GetFromOldID()); err != nil {
			logger.Sugar().Error("FromOldID is invalid: %v", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("FromOldID is invalid: %v", err))
		}
	}

	return nil
}

func duplicate(infos []*npool.DetailReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v", info.AppID, info.UserID, info.CoinTypeID)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:CoinTypeID")
		}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
