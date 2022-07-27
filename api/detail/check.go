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
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.FromCoinTypeID != nil {
		if _, err := uuid.Parse(info.GetFromCoinTypeID()); err != nil {
			logger.Sugar().Errorw("validate", "FromCoinTypeID", info.GetFromCoinTypeID(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("FromCoinTypeID is invalid: %v", err))
		}
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "Amount is less than 0")
		}
	}

	if info.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(info.GetCoinUSDCurrency())
		if err != nil {
			logger.Sugar().Errorw("validate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinUSDCurrency is invalid: %v", err))
		}
		if currency.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "CoinUSDCurrency is less than 0")
		}
	}

	if info.IOType == nil {
		logger.Sugar().Errorw("validate", "IOType", info.IOType)
		return status.Error(codes.InvalidArgument, "IOType is empty")
	}

	if info.IOSubType == nil {
		logger.Sugar().Errorw("validate", "IOSubType", info.IOSubType)
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
			logger.Sugar().Errorw("validate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
			return status.Error(codes.InvalidArgument, "Incoming IOSubType is invalid")
		}
	case npool.IOType_Outcoming:
		switch info.GetIOSubType() {
		case npool.IOSubType_Payment:
		case npool.IOSubType_Withdrawal:
		default:
			logger.Sugar().Errorw("validate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
			return status.Error(codes.InvalidArgument, "Outcoming IOSubType is invalid")
		}
	default:
		logger.Sugar().Errorw("validate", "IOType", info.GetIOType())
		return status.Error(codes.InvalidArgument, "IOSubType is invalid")
	}

	if info.FromOldID != nil {
		if _, err := uuid.Parse(info.GetFromOldID()); err != nil {
			logger.Sugar().Errorw("validate", "FromOldID", info.GetFromOldID(), "error", err)
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

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
