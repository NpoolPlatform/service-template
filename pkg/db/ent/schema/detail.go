package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/message/npool/servicetmpl/detail"
)

// Detail holds the schema definition for the Detail entity.
type Detail struct {
	ent.Schema
}

func (Detail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Detail.
func (Detail) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.UUID("user_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.UUID("coin_type_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.String("io_type").Optional().Default(detail.IOType_DefaultType.String()),
		field.String("io_sub_type").Optional().Default(detail.IOSubType_DefaultSubType.String()),
		field.Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.UUID("from_coin_type_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.Float("coin_usd_currency").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.String("io_extra").Optional().Default(""),
		field.UUID("from_old_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
	}
}

// Edges of the Detail.
func (Detail) Edges() []ent.Edge {
	return nil
}
