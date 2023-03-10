package detail

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/mgr/v1/detail"
)

func trace(span trace1.Span, in *npool.DetailReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("IOType.%v", index), in.GetIOType().String()),
		attribute.String(fmt.Sprintf("IOSubType.%v", index), in.GetIOSubType().String()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.String(fmt.Sprintf("FromCoinTypeID.%v", index), in.GetFromCoinTypeID()),
		attribute.String(fmt.Sprintf("CoinUSDCurrency.%v", index), in.GetCoinUSDCurrency()),
		attribute.String(fmt.Sprintf("IOExtra.%v", index), in.GetIOExtra()),
		attribute.String(fmt.Sprintf("FromOldID.%v", index), in.GetFromOldID()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.DetailReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("IOType.Op", in.GetIOType().GetOp()),
		attribute.String("IOType.Value", npool.IOType(in.GetIOType().GetValue()).String()),
		attribute.String("IOSubType.Op", in.GetIOSubType().GetOp()),
		attribute.String("IOSubType.Value", npool.IOSubType(in.GetIOSubType().GetValue()).String()),
		attribute.String("Amount.Op", in.GetAmount().GetOp()),
		attribute.String("Amount.Value", in.GetAmount().GetValue()),
		attribute.String("FromCoinTypeID.Op", in.GetFromCoinTypeID().GetOp()),
		attribute.String("FromCoinTypeID.Value", in.GetFromCoinTypeID().GetValue()),
		attribute.String("CoinUSDCurrency.Op", in.GetCoinUSDCurrency().GetOp()),
		attribute.String("CoinUSDCurrency.Value", in.GetCoinUSDCurrency().GetValue()),
		attribute.String("IOExtra.Op", in.GetIOExtra().GetOp()),
		attribute.String("IOExtra.Value", in.GetIOExtra().GetValue()),
		attribute.String("FromOldID.Op", in.GetFromOldID().GetOp()),
		attribute.String("FromOldID.Value", in.GetFromOldID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.DetailReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
