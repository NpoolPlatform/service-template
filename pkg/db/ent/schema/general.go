package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// General holds the schema definition for the General entity.
type General struct {
	ent.Schema
}

func (General) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the General.
func (General) Fields() []ent.Field {
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
		field.Float("incoming").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.Float("locked").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.Float("outcoming").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.Float("spendable").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
	}
}

// Edges of the General.
func (General) Edges() []ent.Edge {
	return nil
}
