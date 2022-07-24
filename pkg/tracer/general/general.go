package general

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/servicetmpl/general"
)

func trace(span trace1.Span, in *npool.GeneralReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("Incoming.%v", index), in.GetIncoming()),
		attribute.String(fmt.Sprintf("Locked.%v", index), in.GetLocked()),
		attribute.String(fmt.Sprintf("Outcoming.%v", index), in.GetOutcoming()),
		attribute.String(fmt.Sprintf("Spendable.%v", index), in.GetSpendable()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.GeneralReq) trace1.Span {
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
		attribute.String("Incoming.Op", in.GetIncoming().GetOp()),
		attribute.String("Incoming.Value", in.GetIncoming().GetValue()),
		attribute.String("Locked.Op", in.GetLocked().GetOp()),
		attribute.String("Locked.Value", in.GetLocked().GetValue()),
		attribute.String("Outcoming.Op", in.GetOutcoming().GetOp()),
		attribute.String("Outcoming.Value", in.GetOutcoming().GetValue()),
		attribute.String("Spendable.Op", in.GetSpendable().GetOp()),
		attribute.String("Spendable.Value", in.GetSpendable().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.GeneralReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
