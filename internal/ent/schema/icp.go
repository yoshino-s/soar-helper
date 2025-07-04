package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/yoshino-s/entproto"
)

// Icp holds the schema definition for the Icp entity.
type Icp struct {
	ent.Schema
}

func (Icp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Fields of the Icp.
func (Icp) Fields() []ent.Field {
	return []ent.Field{
		// primary key
		field.String("host").Unique().Immutable().Annotations(entproto.Field(2)),
		field.String("city").Optional().Annotations(entproto.Field(3)),
		field.String("province").Optional().Annotations(entproto.Field(4)),
		field.String("company").Optional().Annotations(entproto.Field(5)),
		field.String("owner").Optional().Annotations(entproto.Field(6)),
		field.String("type").Default("INVALID").Annotations(entproto.Field(7)),
		field.String("homepage").Optional().Annotations(entproto.Field(8)),
		field.String("permit").Optional().Annotations(entproto.Field(9)),
		field.String("webName").Optional().Annotations(entproto.Field(10)),
		field.Time("created_at").Default(time.Now).Annotations(entproto.Field(11)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entproto.Field(12)),
	}
}

// Edges of the Icp.
func (Icp) Edges() []ent.Edge {
	return nil
}
