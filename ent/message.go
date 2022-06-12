// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/message"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// CtreatedAt holds the value of the "ctreated_at" field.
	CtreatedAt time.Time `json:"ctreated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldFrom, message.FieldText:
			values[i] = new(sql.NullString)
		case message.FieldCtreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				m.From = value.String
			}
		case message.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				m.Text = value.String
			}
		case message.FieldCtreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctreated_at", values[i])
			} else if value.Valid {
				m.CtreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", from=")
	builder.WriteString(m.From)
	builder.WriteString(", text=")
	builder.WriteString(m.Text)
	builder.WriteString(", ctreated_at=")
	builder.WriteString(m.CtreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
