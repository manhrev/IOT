// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/manhrev/IOT/server/pkg/ent/data"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
)

// Data is the model entity for the Data schema.
type Data struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Data holds the value of the "data" field.
	Data int64 `json:"data,omitempty"`
	// RecordedAt holds the value of the "recorded_at" field.
	RecordedAt time.Time `json:"recorded_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DataQuery when eager-loading is set.
	Edges     DataEdges `json:"edges"`
	feed_data *int
}

// DataEdges holds the relations/edges for other nodes in the graph.
type DataEdges struct {
	// Feed holds the value of the feed edge.
	Feed *Feed `json:"feed,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FeedOrErr returns the Feed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DataEdges) FeedOrErr() (*Feed, error) {
	if e.loadedTypes[0] {
		if e.Feed == nil {
			// The edge feed was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: feed.Label}
		}
		return e.Feed, nil
	}
	return nil, &NotLoadedError{edge: "feed"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Data) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case data.FieldID, data.FieldData:
			values[i] = new(sql.NullInt64)
		case data.FieldRecordedAt:
			values[i] = new(sql.NullTime)
		case data.ForeignKeys[0]: // feed_data
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Data", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Data fields.
func (d *Data) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case data.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case data.FieldData:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				d.Data = value.Int64
			}
		case data.FieldRecordedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field recorded_at", values[i])
			} else if value.Valid {
				d.RecordedAt = value.Time
			}
		case data.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field feed_data", value)
			} else if value.Valid {
				d.feed_data = new(int)
				*d.feed_data = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryFeed queries the "feed" edge of the Data entity.
func (d *Data) QueryFeed() *FeedQuery {
	return (&DataClient{config: d.config}).QueryFeed(d)
}

// Update returns a builder for updating this Data.
// Note that you need to call Data.Unwrap() before calling this method if this Data
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Data) Update() *DataUpdateOne {
	return (&DataClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Data entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Data) Unwrap() *Data {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Data is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Data) String() string {
	var builder strings.Builder
	builder.WriteString("Data(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", data=")
	builder.WriteString(fmt.Sprintf("%v", d.Data))
	builder.WriteString(", recorded_at=")
	builder.WriteString(d.RecordedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DataSlice is a parsable slice of Data.
type DataSlice []*Data

func (d DataSlice) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
