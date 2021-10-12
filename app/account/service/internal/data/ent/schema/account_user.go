package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type AccountUser struct {
	ent.Schema
}

// Fields of the User.
func (AccountUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Max(11).Comment("账号id"),
		field.String("email").MaxLen(30).Comment("邮箱"),
		field.String("phone").MaxLen(15).Comment("手机号"),
		field.String("username").MaxLen(30).Comment("用户名"),
		field.String("password").MaxLen(32).Comment("密码"),
		field.Int64("created_at").Max(11).Comment("创建时间"),
		field.String("create_ip_at").MaxLen(12).Comment("创建ip"),
		field.Int64("last_login_at").Max(12).Default(0).Comment("最后一次登录时间"),
		field.String("last_login_ip_at").MaxLen(12).Comment("最后一次登录ip"),
		field.Int64("login_times").Max(11).Default(0).Comment("登录次数"),
		field.Int64("status").Max(1).Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
	}
}

// Edges of the User.
func (AccountUser) Edges() []ent.Edge {
	return nil
}
