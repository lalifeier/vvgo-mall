package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AccountUser holds the schema definition for the AccountUser entity.
type AccountUser struct {
	ent.Schema
}

func (AccountUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "account_user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_ci",
		},
	}
}

// Fields of the AccountUser.
func (AccountUser) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.Postgres: "int(11)unsigned", // Override MySQL.
		}).Comment("自增id").Unique(),

		field.String("username").SchemaType(map[string]string{
			dialect.Postgres: "varchar(30)", // Override MySQL.
		}).Default("").Comment("用户名"),

		field.String("phone").SchemaType(map[string]string{
			dialect.Postgres: "varchar(15)", // Override MySQL.
		}).Default("").Comment("手机号"),

		field.String("email").SchemaType(map[string]string{
			dialect.Postgres: "varchar(30)", // Override MySQL.
		}).Default("").Comment("邮箱"),

		field.String("password").SchemaType(map[string]string{
			dialect.Postgres: "varchar(32)", // Override MySQL.
		}).Default("").Comment("密码"),

		field.Int64("create_at").SchemaType(map[string]string{
			dialect.Postgres: "int(11)", // Override MySQL.
		}).Default(0).Comment("创建时间"),

		field.String("create_ip_at").SchemaType(map[string]string{
			dialect.Postgres: "varchar(12)", // Override MySQL.
		}).Default("").Comment("创建ip"),

		field.Int64("last_login_at").SchemaType(map[string]string{
			dialect.Postgres: "int(11)", // Override MySQL.
		}).Default(0).Comment("最后一次登录时间"),

		field.String("last_login_ip_at").SchemaType(map[string]string{
			dialect.Postgres: "varchar(12)", // Override MySQL.
		}).Default("").Comment("最后一次登录ip"),

		field.Int64("login_times").SchemaType(map[string]string{
			dialect.Postgres: "int(11)", // Override MySQL.
		}).Default(0).Comment("登录次数"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.Postgres: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("状态 0:禁用 1:启用 -1:删除"),
	}

}

// Edges of the AccountUser.
func (AccountUser) Edges() []ent.Edge {
	return nil
}
