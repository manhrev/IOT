// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Default: ""},
		{Name: "phone", Type: field.TypeString, Default: ""},
		{Name: "role", Type: field.TypeInt32, Default: 1},
		{Name: "age", Type: field.TypeInt32, Default: 0},
		{Name: "height", Type: field.TypeInt32, Nullable: true},
		{Name: "weight", Type: field.TypeInt32, Nullable: true},
		{Name: "profile_picture", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSettingsColumns holds the columns for the "user_settings" table.
	UserSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "region", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
		{Name: "is_notification", Type: field.TypeString},
		{Name: "date_modified", Type: field.TypeTime},
		{Name: "user_user_setting", Type: field.TypeInt64, Unique: true},
	}
	// UserSettingsTable holds the schema information for the "user_settings" table.
	UserSettingsTable = &schema.Table{
		Name:       "user_settings",
		Columns:    UserSettingsColumns,
		PrimaryKey: []*schema.Column{UserSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_settings_users_user_setting",
				Columns:    []*schema.Column{UserSettingsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserSettingsTable,
	}
)

func init() {
	UserSettingsTable.ForeignKeys[0].RefTable = UsersTable
}