package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"github.com/lalifeier/vvgo-mall/pkg/util/entgo/mixin"
)

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
