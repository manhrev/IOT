package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Data struct {
	ent.Schema
}

func (Data) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("data"),
		field.Time("recorded_at").Default(time.Now()),
	}
}

func (Data) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feed", Feed.Type).Ref("data").Unique(),
	}
}
