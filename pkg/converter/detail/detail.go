package detail

import (
	npool "github.com/NpoolPlatform/message/npool/servicetmpl/detail"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"
)

func Ent2Grpc(row *ent.Detail) *npool.Detail {
	return &npool.Detail{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		IOType:          npool.IOType(npool.IOType_value[row.IoType]),
		IOSubType:       npool.IOSubType(npool.IOSubType_value[row.IoSubType]),
		Amount:          row.Amount.String(),
		FromCoinTypeID:  row.FromCoinTypeID.String(),
		CoinUSDCurrency: row.CoinUsdCurrency.String(),
		IOExtra:         row.IoExtra,
		FromOldID:       row.FromOldID.String(),
	}
}

func Ent2GrpcMany(rows []*ent.Detail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
