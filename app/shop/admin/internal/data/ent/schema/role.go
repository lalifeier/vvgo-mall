package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/pkg/ent/mixin"
)

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_role",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("角色id").Unique(),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.

		}).Default("").Comment("角色名称"),

		field.Int8("sort").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).Default(0).Comment("排序"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用"),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("备注"),
	}

}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
