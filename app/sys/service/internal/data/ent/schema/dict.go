package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo/pkg/ent/mixin"
)

// Dict holds the schema definition for the Dict entity.
type Dict struct {
	ent.Schema
}

func (Dict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

func (Dict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Dict.
func (Dict) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Comment("字典数据id").Unique(),

		// field.Int64("dict_type_id").SchemaType(map[string]string{
		// 	dialect.MySQL: "intunsigned", // Override MySQL.
		// }).Default(0).Comment("字典类型id"),

		field.String("type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("字典类型"),

		field.String("label").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("字典标签"),

		field.String("value").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("字典键值"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1) unsigned", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用 -1:删除"),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("备注"),

		field.Int8("sort").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1) unsigned", // Override MySQL.
		}).Default(0).Comment("排序"),

		field.Int8("is_default").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1) unsigned", // Override MySQL.
		}).Default(0).Comment("是否默认值 0:否 1:是"),
	}
}

// Edges of the Dict.
func (Dict) Edges() []ent.Edge {
	return nil
}
