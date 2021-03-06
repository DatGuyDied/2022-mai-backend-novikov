// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/message"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/schema"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/usermessage"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCtreatedAt is the schema descriptor for ctreated_at field.
	messageDescCtreatedAt := messageFields[2].Descriptor()
	// message.DefaultCtreatedAt holds the default value on creation for the ctreated_at field.
	message.DefaultCtreatedAt = messageDescCtreatedAt.Default.(func() time.Time)
	usermessageFields := schema.UserMessage{}.Fields()
	_ = usermessageFields
	// usermessageDescCtreatedAt is the schema descriptor for ctreated_at field.
	usermessageDescCtreatedAt := usermessageFields[3].Descriptor()
	// usermessage.DefaultCtreatedAt holds the default value on creation for the ctreated_at field.
	usermessage.DefaultCtreatedAt = usermessageDescCtreatedAt.Default.(func() time.Time)
}
