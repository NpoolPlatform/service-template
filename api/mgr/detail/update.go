//nolint:nolintlint,dupl
package detail

import (
	"context"
	"fmt"

	commontracer "github.com/NpoolPlatform/service-template/pkg/tracer"

	converter "github.com/NpoolPlatform/service-template/pkg/mgr/detail/converter"
	crud "github.com/NpoolPlatform/service-template/pkg/mgr/detail/crud"
	tracer "github.com/NpoolPlatform/service-template/pkg/mgr/detail/tracer"

	constant "github.com/NpoolPlatform/service-template/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func ValidateUpdate(info *npool.DetailReq) error { //nolint
	if _, err := uuid.Parse(info.GetID()); err != nil {
		logger.Sugar().Errorw("ValidateUpdate", "ID", info.GetID(), "Error", err)
		return err
	}

	if info.AppID != nil {
		if _, err := uuid.Parse(info.GetAppID()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "AppID", info.GetAppID(), "Error", err)
			return err
		}
	}

	if info.UserID != nil {
		if _, err := uuid.Parse(info.GetUserID()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "UserID", info.GetUserID(), "Error", err)
			return err
		}
	}

	if info.CoinTypeID != nil {
		if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "CoinTypeID", info.GetCoinTypeID(), "Error", err)
			return err
		}
	}

	if info.FromCoinTypeID != nil {
		if _, err := uuid.Parse(info.GetFromCoinTypeID()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "FromCoinTypeID", info.GetFromCoinTypeID(), "Error", err)
			return err
		}
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "Amount", info.GetAmount(), "Error", err)
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("ValidateUpdate", "Amount", info.GetAmount(), "Error", "Amount < 0")
			return fmt.Errorf("amount < 0")
		}
	}

	if info.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(info.GetCoinUSDCurrency())
		if err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "Error", err)
			return err
		}
		if currency.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("ValidateUpdate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "Error", "CoinUSDCurrency < 0")
			return fmt.Errorf("coinusdcurrency < 0")
		}
	}

	if info.IOType != nil && info.IOSubType != nil {
		switch info.GetIOType() {
		case npool.IOType_Incoming:
			switch info.GetIOSubType() {
			case npool.IOSubType_Payment:
			case npool.IOSubType_MiningBenefit:
			case npool.IOSubType_Commission:
			case npool.IOSubType_TechniqueFeeCommission:
			default:
				logger.Sugar().Errorw("ValidateUpdate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
				return fmt.Errorf("incoming iosubtype invalid")
			}
		case npool.IOType_Outcoming:
			switch info.GetIOSubType() {
			case npool.IOSubType_Payment:
			case npool.IOSubType_Withdrawal:
			default:
				logger.Sugar().Errorw("ValidateUpdate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
				return fmt.Errorf("outcoming iosubtype invalid")
			}
		default:
			logger.Sugar().Errorw("ValidateUpdate", "IOType", info.GetIOType())
			return fmt.Errorf("iosubtype invalid")
		}
	}

	if info.FromOldID != nil {
		if _, err := uuid.Parse(info.GetFromOldID()); err != nil {
			logger.Sugar().Errorw("ValidateUpdate", "FromOldID", info.GetFromOldID(), "error", err)
			return err
		}
	}

	return nil
}

func (s *Server) UpdateDetail(ctx context.Context, in *npool.UpdateDetailRequest) (*npool.UpdateDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = ValidateUpdate(in.GetInfo())
	if err != nil {
		return &npool.UpdateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "detail", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateDetail", "Info", in.GetInfo(), "Error", err)
		return &npool.UpdateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
