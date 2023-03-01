package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/pkg/util/entgo/mixin"
)

func (Api) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_api",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

func (Api) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Api holds the schema definition for the Api entity.
type Api struct {
	ent.Schema
}

// Fields of the Api.
func (Api) Fields() []ent.Field {

	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("自增id").Unique(),

		field.String("group").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("接口分组"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Comment("接口名称"),

		field.String("path").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("接口路径"),

		field.String("method").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}).Default("").Comment("接口请求方式"),

		field.String("desc").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("接口描述"),

		field.String("permission").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("接口权限"),

		field.Uint8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用 -1:删除"),
	}

}

// Edges of the Api.
func (Api) Edges() []ent.Edge {
	return nil
}
