package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Feed struct {
	ent.Schema
}

func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.String("feed_name").Unique(),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("data", Data.Type),
	}
}
