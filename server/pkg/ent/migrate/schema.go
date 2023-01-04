// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DataColumns holds the columns for the "data" table.
	DataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "data", Type: field.TypeInt64},
		{Name: "recorded_at", Type: field.TypeTime},
		{Name: "feed_data", Type: field.TypeInt, Nullable: true},
	}
	// DataTable holds the schema information for the "data" table.
	DataTable = &schema.Table{
		Name:       "data",
		Columns:    DataColumns,
		PrimaryKey: []*schema.Column{DataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "data_feeds_data",
				Columns:    []*schema.Column{DataColumns[3]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "feed_name", Type: field.TypeString, Unique: true},
		{Name: "data_type", Type: field.TypeUint16, Default: 0},
		{Name: "display_type", Type: field.TypeUint16, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:       "feeds",
		Columns:    FeedsColumns,
		PrimaryKey: []*schema.Column{FeedsColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group_name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// GroupFeedsColumns holds the columns for the "group_feeds" table.
	GroupFeedsColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "feed_id", Type: field.TypeInt},
	}
	// GroupFeedsTable holds the schema information for the "group_feeds" table.
	GroupFeedsTable = &schema.Table{
		Name:       "group_feeds",
		Columns:    GroupFeedsColumns,
		PrimaryKey: []*schema.Column{GroupFeedsColumns[0], GroupFeedsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_feeds_group_id",
				Columns:    []*schema.Column{GroupFeedsColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_feeds_feed_id",
				Columns:    []*schema.Column{GroupFeedsColumns[1]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DataTable,
		FeedsTable,
		GroupsTable,
		GroupFeedsTable,
	}
)

func init() {
	DataTable.ForeignKeys[0].RefTable = FeedsTable
	GroupFeedsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupFeedsTable.ForeignKeys[1].RefTable = FeedsTable
}
