package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"time"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

func (Staff) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "staff",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("自增id").Unique(),

		field.Int64("uid").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("账号id"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}).Default("").Comment("员工姓名"),

		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(15)", // Override MySQL.
		}).Default("").Comment("员工手机号"),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}).Default("").Comment("员工邮箱"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}).Default("").Comment("员工昵称"),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("员工头像(相对路径)"),

		// field.Int8("gender").SchemaType(map[string]string{
		// 	dialect.MySQL: "tinyint(1) unsigned", // Override MySQL.
		// }).Default(0).Comment("员工性别 1:男性 2:女性"),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).Comment("创建时间"),

		field.Int64("created_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("更新人"),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),

		field.Int64("updated_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("更新人"),
	}

}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return nil
}
