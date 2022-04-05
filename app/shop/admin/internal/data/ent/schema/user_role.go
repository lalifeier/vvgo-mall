package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_user_role",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("用户ID"),

		field.Int64("role_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("角色ID"),
	}

}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}

func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").
			Unique().StorageKey("idx_user_id_role_id"),
	}
}
