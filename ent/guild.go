// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/db/ent/guild"
	"github.com/roleypoly/db/ent/schema"
)

// Guild is the model entity for the Guild schema.
type Guild struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Snowflake holds the value of the "snowflake" field.
	Snowflake string `json:"snowflake,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Categories holds the value of the "categories" field.
	Categories []schema.Category `json:"categories,omitempty"`
	// Entitlements holds the value of the "entitlements" field.
	Entitlements []string `json:"entitlements,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Guild) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // snowflake
		&sql.NullString{}, // message
		&[]byte{},         // categories
		&[]byte{},         // entitlements
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Guild fields.
func (gu *Guild) assignValues(values ...interface{}) error {
	if m, n := len(values), len(guild.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gu.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		gu.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		gu.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field snowflake", values[2])
	} else if value.Valid {
		gu.Snowflake = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field message", values[3])
	} else if value.Valid {
		gu.Message = value.String
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field categories", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &gu.Categories); err != nil {
			return fmt.Errorf("unmarshal field categories: %v", err)
		}
	}

	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field entitlements", values[5])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &gu.Entitlements); err != nil {
			return fmt.Errorf("unmarshal field entitlements: %v", err)
		}
	}
	return nil
}

// Update returns a builder for updating this Guild.
// Note that, you need to call Guild.Unwrap() before calling this method, if this Guild
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guild) Update() *GuildUpdateOne {
	return (&GuildClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gu *Guild) Unwrap() *Guild {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Guild is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guild) String() string {
	var builder strings.Builder
	builder.WriteString("Guild(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(gu.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(gu.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", snowflake=")
	builder.WriteString(gu.Snowflake)
	builder.WriteString(", message=")
	builder.WriteString(gu.Message)
	builder.WriteString(", categories=")
	builder.WriteString(fmt.Sprintf("%v", gu.Categories))
	builder.WriteString(", entitlements=")
	builder.WriteString(fmt.Sprintf("%v", gu.Entitlements))
	builder.WriteByte(')')
	return builder.String()
}

// Guilds is a parsable slice of Guild.
type Guilds []*Guild

func (gu Guilds) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
