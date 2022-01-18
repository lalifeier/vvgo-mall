package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lalifeier/vvgo/pkg/ent/hooks"
)

// AuditMixin 实现了 ent.Mixin，
// 用于 schema 包内通用的审计日志功能。
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).Comment("创建时间"),

		field.Int64("created_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("更新人"),

		field.Time("updated_at").Immutable().SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),

		field.Int64("updated_by").SchemaType(map[string]string{
			dialect.MySQL: "int(11)unsigned", // Override MySQL.
		}).Default(0).Comment("更新人"),
	}
}

// AuditMixin 的钩子
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.AuditHook,
	}
}
