package detail

/*
import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mw/v1/detail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreate(info *npool.DetailReq) error { //nolint
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "AppID", info.GetAppID(), "Error", err)
		return err
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "UserID", info.GetUserID(), "Error", err)
		return err
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("ValidateCreate", "CoinTypeID", info.GetCoinTypeID(), "Error", err)
		return err
	}

	if info.FromCoinTypeID != nil {
		if _, err := uuid.Parse(info.GetFromCoinTypeID()); err != nil {
			logger.Sugar().Errorw("ValidateCreate", "FromCoinTypeID", info.GetFromCoinTypeID(), "Error", err)
			return err
		}
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("ValidateCreate", "Amount", info.GetAmount(), "Error", err)
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("ValidateCreate", "Amount", info.GetAmount(), "Error", "Amount < 0")
			return fmt.Errorf("amount < 0")
		}
	}

	if info.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(info.GetCoinUSDCurrency())
		if err != nil {
			logger.Sugar().Errorw("ValidateCreate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "Error", err)
			return err
		}
		if currency.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("ValidateCreate", "CoinUSDCurrency", info.GetCoinUSDCurrency(), "Error", "CoinUSDCurrency < 0")
			return fmt.Errorf("coinusdcurrency < 0")
		}
	}

	switch info.GetIOType() {
	case npool.IOType_Incoming:
		switch info.GetIOSubType() {
		case npool.IOSubType_Payment:
		case npool.IOSubType_MiningBenefit:
		case npool.IOSubType_Commission:
		case npool.IOSubType_TechniqueFeeCommission:
		default:
			logger.Sugar().Errorw("ValidateCreate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
			return fmt.Errorf("incoming iosubtype invalid")
		}
	case npool.IOType_Outcoming:
		switch info.GetIOSubType() {
		case npool.IOSubType_Payment:
		case npool.IOSubType_Withdrawal:
		default:
			logger.Sugar().Errorw("ValidateCreate", "IOType", info.GetIOType(), "IOSubType", info.GetIOSubType())
			return fmt.Errorf("outcoming iosubtype invalid")
		}
	default:
		logger.Sugar().Errorw("ValidateCreate", "IOType", info.GetIOType())
		return fmt.Errorf("iosubtype invalid")
	}

	if info.FromOldID != nil {
		if _, err := uuid.Parse(info.GetFromOldID()); err != nil {
			logger.Sugar().Errorw("ValidateCreate", "FromOldID", info.GetFromOldID(), "error", err)
			return err
		}
	}

	return nil
}

func ValidateCreateMany(in []*npool.DetailReq) error {
	for _, info := range in {
		if err := ValidateCreate(info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) CreateDetail(ctx context.Context, in *npool.CreateDetailRequest) (*npool.CreateDetailResponse, error) {

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateDetail", "Info", in.GetInfo(), "Error", err)
		return &npool.CreateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateDetails(ctx context.Context, in *npool.CreateDetailsRequest) (*npool.CreateDetailsResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateDetailsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = ValidateCreateMany(in.GetInfos())
	if err != nil {
		return &npool.CreateDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateDetail", "Infos", len(in.GetInfos()), "Error", err)
		return &npool.CreateDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}
*/
