package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// AuditMixin 实现了 ent.Mixin，
// 用于 schema 包内通用的审计日志功能。
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Int("created_by").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by").
			Optional(),
	}
}

// // AuditMixin 的钩子
// func (AuditMixin) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hooks.AuditHook,
// 	}
// }

// // AuditHook 是审计日志钩子的示例
// func AuditHook(next ent.Mutator) ent.Mutator {
// 	// AuditLogger 声明了所有嵌入 AuditLog mixin 的
// 	// Schema 更变所共有的方法。 若变量 "existence "为真，
// 	// 则该字段存在于此更变中 (例如被其它的钩子设置过)。
// 	type AuditLogger interface {
// 		SetCreatedAt(time.Time)
// 		CreatedAt() (value time.Time, exists bool)
// 		SetCreatedBy(int)
// 		CreatedBy() (id int, exists bool)
// 		SetUpdatedAt(time.Time)
// 		UpdatedAt() (value time.Time, exists bool)
// 		SetUpdatedBy(int)
// 		UpdatedBy() (id int, exists bool)
// 	}
// 	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 		ml, ok := m.(AuditLogger)
// 		if !ok {
// 			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
// 		}
// 		usr, err := viewer.UserFromContext(ctx)
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch op := m.Op(); {
// 		case op.Is(ent.OpCreate):
// 			ml.SetCreatedAt(time.Now())
// 			if _, exists := ml.CreatedBy(); !exists {
// 				ml.SetCreatedBy(usr.ID)
// 			}
// 		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
// 			ml.SetUpdatedAt(time.Now())
// 			if _, exists := ml.UpdatedBy(); !exists {
// 				ml.SetUpdatedBy(usr.ID)
// 			}
// 		}
// 		return next.Mutate(ctx, m)
// 	})
// }
