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
		field.String("city").Default(""),
		field.String("province").Default(""),
		field.String("company"),
		field.String("owner"),
		field.String("type").Default("INVALID"),
		field.String("homepage"),
		field.String("permit"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
	}
}

// Edges of the Icp.
func (Icp) Edges() []ent.Edge {
	return nil
}
