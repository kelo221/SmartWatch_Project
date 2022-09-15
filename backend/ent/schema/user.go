package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.Time("createdAt").
			Default(time.Now()),
		field.Bool("elevated").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
	}
}
