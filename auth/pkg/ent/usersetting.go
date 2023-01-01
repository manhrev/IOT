// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/user"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/usersetting"
)

// UserSetting is the model entity for the UserSetting schema.
type UserSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// IsNotification holds the value of the "is_notification" field.
	IsNotification string `json:"is_notification,omitempty"`
	// DateModified holds the value of the "date_modified" field.
	DateModified time.Time `json:"date_modified,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSettingQuery when eager-loading is set.
	Edges             UserSettingEdges `json:"edges"`
	user_user_setting *int64
}

// UserSettingEdges holds the relations/edges for other nodes in the graph.
type UserSettingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSettingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersetting.FieldID:
			values[i] = new(sql.NullInt64)
		case usersetting.FieldRegion, usersetting.FieldLanguage, usersetting.FieldIsNotification:
			values[i] = new(sql.NullString)
		case usersetting.FieldDateModified:
			values[i] = new(sql.NullTime)
		case usersetting.ForeignKeys[0]: // user_user_setting
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSetting fields.
func (us *UserSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int64(value.Int64)
		case usersetting.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				us.Region = value.String
			}
		case usersetting.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				us.Language = value.String
			}
		case usersetting.FieldIsNotification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_notification", values[i])
			} else if value.Valid {
				us.IsNotification = value.String
			}
		case usersetting.FieldDateModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_modified", values[i])
			} else if value.Valid {
				us.DateModified = value.Time
			}
		case usersetting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_user_setting", value)
			} else if value.Valid {
				us.user_user_setting = new(int64)
				*us.user_user_setting = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserSetting entity.
func (us *UserSetting) QueryUser() *UserQuery {
	return (&UserSettingClient{config: us.config}).QueryUser(us)
}

// Update returns a builder for updating this UserSetting.
// Note that you need to call UserSetting.Unwrap() before calling this method if this UserSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSetting) Update() *UserSettingUpdateOne {
	return (&UserSettingClient{config: us.config}).UpdateOne(us)
}

// Unwrap unwraps the UserSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSetting) Unwrap() *UserSetting {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSetting is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSetting) String() string {
	var builder strings.Builder
	builder.WriteString("UserSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("region=")
	builder.WriteString(us.Region)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(us.Language)
	builder.WriteString(", ")
	builder.WriteString("is_notification=")
	builder.WriteString(us.IsNotification)
	builder.WriteString(", ")
	builder.WriteString("date_modified=")
	builder.WriteString(us.DateModified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserSettings is a parsable slice of UserSetting.
type UserSettings []*UserSetting

func (us UserSettings) config(cfg config) {
	for _i := range us {
		us[_i].config = cfg
	}
}