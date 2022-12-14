// Code generated by entc, DO NOT EDIT.

package group

import (
	"time"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupName holds the string denoting the group_name field in the database.
	FieldGroupName = "group_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeFeeds holds the string denoting the feeds edge name in mutations.
	EdgeFeeds = "feeds"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// FeedsTable is the table that holds the feeds relation/edge. The primary key declared below.
	FeedsTable = "group_feeds"
	// FeedsInverseTable is the table name for the Feed entity.
	// It exists in this package in order to avoid circular dependency with the "feed" package.
	FeedsInverseTable = "feeds"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldGroupName,
	FieldCreatedAt,
}

var (
	// FeedsPrimaryKey and FeedsColumn2 are the table columns denoting the
	// primary key for the feeds relation (M2M).
	FeedsPrimaryKey = []string{"group_id", "feed_id"}
)

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
