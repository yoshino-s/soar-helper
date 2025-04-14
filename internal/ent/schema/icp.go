package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Icp holds the schema definition for the Icp entity.
type Icp struct {
	ent.Schema
}

// Fields of the Icp.
func (Icp) Fields() []ent.Field {
	return []ent.Field{
		// primary key
		field.String("host").Unique().Immutable(),
		field.String("city").Optional(),
		field.String("province").Optional(),
		field.String("company").Optional(),
		field.String("owner").Optional(),
		field.String("type").Default("INVALID"),
		field.String("homepage").Optional(),
		field.String("permit").Optional(),
		field.String("webName").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Icp.
func (Icp) Edges() []ent.Edge {
	return nil
}
