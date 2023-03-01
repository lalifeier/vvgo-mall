package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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

	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("用户id").Unique(),

		field.Int64("uid").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("账号id"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}).Default("").Comment("昵称"),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("头像"),

		field.Enum("gender").SchemaType(map[string]string{
			dialect.MySQL: "enum('male','female','unknow')", // Override MySQL.
		}).Default("unknow").Values("male", "female", "unknow").Comment("性别"),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("备注"),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
