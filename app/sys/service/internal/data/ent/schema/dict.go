package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Dict holds the schema definition for the Dict entity.
type Dict struct {
	ent.Schema
}

// Fields of the Dict.
func (Dict) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigintunsigned", // Override MySQL.
		}).Comment("字典数据id").Unique(),

		field.Int64("dict_type_id").SchemaType(map[string]string{
			dialect.MySQL: "bigintunsigned", // Override MySQL.
		}).Comment("字典类型id"),

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
			dialect.MySQL: "tinyint(1)unsigned", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用"),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(500)", // Override MySQL.
		}).Default("").Comment("备注"),

		field.Int8("sort").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)unsigned", // Override MySQL.
		}).Default(0).Comment("排序"),

		field.Int8("is_default").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)unsigned", // Override MySQL.
		}).Default(0).Comment("是否默认值 0:否 1:是"),

		field.Int32("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("创建时间"),

		field.Int32("create_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("更新人"),

		field.Int32("update_at").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("更新时间"),

		field.Int32("update_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("更新人"),
	}

}

// Edges of the Dict.
func (Dict) Edges() []ent.Edge {
	return nil
}
