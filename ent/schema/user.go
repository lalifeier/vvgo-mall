// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.Int64("id"), field.Int64("uid"), field.String("name"), field.String("phone"), field.String("email"), field.String("nickname"), field.String("avatar"), field.String("gender"), field.String("remark"), field.Int64("dept_id"), field.Int64("created_at"), field.Int64("created_by"), field.Int64("updated_at"), field.Int64("updated_by")}
}
func (User) Edges() []ent.Edge {
	return nil
}
func (User) Annotations() []schema.Annotation {
	return nil
}
