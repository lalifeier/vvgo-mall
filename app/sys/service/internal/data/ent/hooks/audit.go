package hooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
)

// // AuditHook 是审计日志钩子的示例
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger 声明了所有嵌入 AuditLog mixin 的
	// Schema 更变所共有的方法。 若变量 "existence "为真，
	// 则该字段存在于此更变中 (例如被其它的钩子设置过)。
	type AuditLogger interface {
		SetCreateAt(time.Time)
		CreateAt() (value time.Time, exists bool)
		SetCreateBy(int64)
		CreateBy() (id int64, exists bool)
		SetUpdateAt(time.Time)
		UpdateAt() (value time.Time, exists bool)
		SetUpdateBy(int64)
		UpdateBy() (id int64, exists bool)
	}

	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		// user, err := viewer.UserFromContext(ctx)
		// if err != nil {
		// 	return nil, err
		// }
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreateAt(time.Now())
			if _, exists := ml.CreateBy(); !exists {
				// 		// ml.SetCreateBy(user.ID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdateAt(time.Now())
			if _, exists := ml.UpdateBy(); !exists {
				// 		// ml.SetUpdateBy(user.ID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
