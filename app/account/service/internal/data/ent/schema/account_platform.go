package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// AccountPlatform holds the schema definition for the AccountPlatform entity.
type AccountPlatform struct {
	ent.Schema
}

// Fields of the AccountPlatform.
func (AccountPlatform) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Comment("自增id").Unique(),

		field.Int32("uid").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("账号id"),

		field.String("platform_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(60)", // Override MySQL.
		}).Default("").Comment("平台id"),

		field.String("platform_token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(60)", // Override MySQL.
		}).Default("").Comment("平台access_token"),

		field.Int8("type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("平台类型 0:未知,1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(60)", // Override MySQL.
		}).Default("").Comment("昵称"),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Comment("头像"),

		field.Int32("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("创建时间"),

		field.Int32("update_at").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("更新时间"),
	}

}

// Edges of the AccountPlatform.
func (AccountPlatform) Edges() []ent.Edge {
	return nil
}
