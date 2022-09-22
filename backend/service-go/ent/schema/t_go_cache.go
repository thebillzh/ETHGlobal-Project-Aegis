// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type TGoCache struct {
	ent.Schema
}

func (TGoCache) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("auto increment primary key"),
		field.String("cache_key").Unique().Comment("cache key"),
		field.String("cache_value").Default("").Comment("cached value"),
		field.Time("mtime").Comment("modify time").Default(time.Now()).UpdateDefault(time.Now),
		field.Time("ctime").Comment("create time").Default(time.Now()),
	}
}
func (TGoCache) Edges() []ent.Edge {
	return nil
}
func (TGoCache) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_go_caches"},
	}
}
