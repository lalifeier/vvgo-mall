package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type AccountPlatform struct {
	ent.Schema
}

// Fields of the User.
func (AccountPlatform) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Max(11).Comment("自增id"),
		field.Int64("uid").Max(11).Default(0).Comment("账号id"),
		field.String("platform_id").MaxLen(60).Comment("平台id"),
		field.String("platform_token").MaxLen(60).Comment("平台access_token"),
		field.Int64("type").Max(1).Default(0).Comment("平台类型 0:未知,1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter"),
		field.String("nickname").MaxLen(60).Comment("昵称"),
		field.String("avatar").MaxLen(255).Comment("头像"),
		field.Int64("create_at").Max(11).Comment("创建时间"),
		field.Int64("update_at").Max(11).Comment("更新时间"),
	}
}

// Edges of the User.
func (AccountPlatform) Edges() []ent.Edge {
	return nil
}
