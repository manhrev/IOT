// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"
)

const (
	// Label holds the string label denoting the feed type in the database.
	Label = "feed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFeedName holds the string denoting the feed_name field in the database.
	FieldFeedName = "feed_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeData holds the string denoting the data edge name in mutations.
	EdgeData = "data"
	// Table holds the table name of the feed in the database.
	Table = "feeds"
	// DataTable is the table that holds the data relation/edge.
	DataTable = "data"
	// DataInverseTable is the table name for the Data entity.
	// It exists in this package in order to avoid circular dependency with the "data" package.
	DataInverseTable = "data"
	// DataColumn is the table column denoting the data relation/edge.
	DataColumn = "feed_data"
)

// Columns holds all SQL columns for feed fields.
var Columns = []string{
	FieldID,
	FieldFeedName,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)
