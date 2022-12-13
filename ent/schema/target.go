package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Target struct {
	ent.Schema
}

func (Target) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.String("name"),
		field.Enum("type").
			Values("TYPE_HTTP", "TYPE_TCP", "TYPE_PING"),
		field.Uint32("interval_seconds").
			Positive(),
		field.Uint32("timeout_seconds").
			Positive(),
		field.String("url").
			Unique().
			Optional().
			Nillable(),
		field.String("hostname").
			Optional().
			Nillable(),
		field.Uint32("port").
			Positive().
			Optional().
			Nillable(),
	}
}

func (Target) Edges() []ent.Edge {
	return nil
}

func (Target) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "type").Unique(),
	}
}
