package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

type Feed struct {
	ent.Schema
}

func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.String("feed_name").Unique(),
		field.Uint16("data_type").Default(uint16(info_pb.FeedDataType_FEED_DATA_TYPE_UNSPECIFIED)),
		field.Uint16("display_type").Default(uint16(info_pb.FeedDisplayType_FEED_TYPE_UNSPECIFIED)),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("data", Data.Type),
		edge.From("group", Group.Type).Ref("feeds"),
	}
}
