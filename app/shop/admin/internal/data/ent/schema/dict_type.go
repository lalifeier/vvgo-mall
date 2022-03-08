package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/pkg/ent/mixin"
)

func (DictType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "sys_dict_type",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

func (DictType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// DictType holds the schema definition for the DictType entity.
type DictType struct {
	ent.Schema
}

// Fields of the DictType.
func (DictType) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("字典类型id").Unique(),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("字典名称"),

		field.String("type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Default("").Comment("字典类型"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)unsigned", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用"),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("备注"),
	}

}

// Edges of the DictType.
func (DictType) Edges() []ent.Edge {
	return nil
}
