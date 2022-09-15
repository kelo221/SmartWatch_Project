package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Int("unixTime").
			Default(int(time.Now().Unix()) + 60).
			NonNegative(),
		field.String("eventName").
			Default("unknown"),
		field.Bool("isSilent").
			Default(false),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("username", User.Type).
			Ref("events").
			Unique(),
	}
}

// Indexes of the Event.
func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("eventName").
			Edges("username").
			Unique(),
	}
}
