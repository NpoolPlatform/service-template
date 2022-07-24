package general

import (
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/general"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
)

func Ent2Grpc(row *ent.General) *npool.General {
	return &npool.General{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Incoming:   row.Incoming.String(),
		Locked:     row.Locked.String(),
		Outcoming:  row.Outcoming.String(),
		Spendable:  row.Spendable.String(),
	}
}

func Ent2GrpcMany(rows []*ent.General) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
