package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_name").Unique(),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feeds", Feed.Type),
	}
}
