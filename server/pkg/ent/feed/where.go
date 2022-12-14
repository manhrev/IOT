// Code generated by entc, DO NOT EDIT.

package feed

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/manhrev/IOT/server/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FeedName applies equality check predicate on the "feed_name" field. It's identical to FeedNameEQ.
func FeedName(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedName), v))
	})
}

// DataType applies equality check predicate on the "data_type" field. It's identical to DataTypeEQ.
func DataType(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataType), v))
	})
}

// DisplayType applies equality check predicate on the "display_type" field. It's identical to DisplayTypeEQ.
func DisplayType(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayType), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// FeedNameEQ applies the EQ predicate on the "feed_name" field.
func FeedNameEQ(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedName), v))
	})
}

// FeedNameNEQ applies the NEQ predicate on the "feed_name" field.
func FeedNameNEQ(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeedName), v))
	})
}

// FeedNameIn applies the In predicate on the "feed_name" field.
func FeedNameIn(vs ...string) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFeedName), v...))
	})
}

// FeedNameNotIn applies the NotIn predicate on the "feed_name" field.
func FeedNameNotIn(vs ...string) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFeedName), v...))
	})
}

// FeedNameGT applies the GT predicate on the "feed_name" field.
func FeedNameGT(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeedName), v))
	})
}

// FeedNameGTE applies the GTE predicate on the "feed_name" field.
func FeedNameGTE(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeedName), v))
	})
}

// FeedNameLT applies the LT predicate on the "feed_name" field.
func FeedNameLT(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeedName), v))
	})
}

// FeedNameLTE applies the LTE predicate on the "feed_name" field.
func FeedNameLTE(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeedName), v))
	})
}

// FeedNameContains applies the Contains predicate on the "feed_name" field.
func FeedNameContains(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFeedName), v))
	})
}

// FeedNameHasPrefix applies the HasPrefix predicate on the "feed_name" field.
func FeedNameHasPrefix(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFeedName), v))
	})
}

// FeedNameHasSuffix applies the HasSuffix predicate on the "feed_name" field.
func FeedNameHasSuffix(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFeedName), v))
	})
}

// FeedNameEqualFold applies the EqualFold predicate on the "feed_name" field.
func FeedNameEqualFold(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFeedName), v))
	})
}

// FeedNameContainsFold applies the ContainsFold predicate on the "feed_name" field.
func FeedNameContainsFold(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFeedName), v))
	})
}

// DataTypeEQ applies the EQ predicate on the "data_type" field.
func DataTypeEQ(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataType), v))
	})
}

// DataTypeNEQ applies the NEQ predicate on the "data_type" field.
func DataTypeNEQ(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDataType), v))
	})
}

// DataTypeIn applies the In predicate on the "data_type" field.
func DataTypeIn(vs ...uint16) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDataType), v...))
	})
}

// DataTypeNotIn applies the NotIn predicate on the "data_type" field.
func DataTypeNotIn(vs ...uint16) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDataType), v...))
	})
}

// DataTypeGT applies the GT predicate on the "data_type" field.
func DataTypeGT(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDataType), v))
	})
}

// DataTypeGTE applies the GTE predicate on the "data_type" field.
func DataTypeGTE(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDataType), v))
	})
}

// DataTypeLT applies the LT predicate on the "data_type" field.
func DataTypeLT(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDataType), v))
	})
}

// DataTypeLTE applies the LTE predicate on the "data_type" field.
func DataTypeLTE(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDataType), v))
	})
}

// DisplayTypeEQ applies the EQ predicate on the "display_type" field.
func DisplayTypeEQ(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayType), v))
	})
}

// DisplayTypeNEQ applies the NEQ predicate on the "display_type" field.
func DisplayTypeNEQ(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplayType), v))
	})
}

// DisplayTypeIn applies the In predicate on the "display_type" field.
func DisplayTypeIn(vs ...uint16) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDisplayType), v...))
	})
}

// DisplayTypeNotIn applies the NotIn predicate on the "display_type" field.
func DisplayTypeNotIn(vs ...uint16) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDisplayType), v...))
	})
}

// DisplayTypeGT applies the GT predicate on the "display_type" field.
func DisplayTypeGT(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplayType), v))
	})
}

// DisplayTypeGTE applies the GTE predicate on the "display_type" field.
func DisplayTypeGTE(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplayType), v))
	})
}

// DisplayTypeLT applies the LT predicate on the "display_type" field.
func DisplayTypeLT(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplayType), v))
	})
}

// DisplayTypeLTE applies the LTE predicate on the "display_type" field.
func DisplayTypeLTE(v uint16) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplayType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasData applies the HasEdge predicate on the "data" edge.
func HasData() predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DataTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DataTable, DataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDataWith applies the HasEdge predicate on the "data" edge with a given conditions (other predicates).
func HasDataWith(preds ...predicate.Data) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DataTable, DataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupTable, GroupPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupTable, GroupPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		p(s.Not())
	})
}
