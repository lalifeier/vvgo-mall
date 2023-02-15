package hooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
)

// AuditHook 是审计日志钩子的示例
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger 声明了所有嵌入 AuditLog mixin 的
	// Schema 更变所共有的方法。 若变量 "existence "为真，
	// 则该字段存在于此更变中 (例如被其它的钩子设置过)。
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(int)
		CreatedBy() (id int, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(int)
		UpdatedBy() (id int, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		// usr, err := viewer.UserFromContext(ctx)
		// if err != nil {
		// 		return nil, err
		// }
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				// ml.SetCreatedBy(usr.ID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				// ml.SetUpdatedBy(usr.ID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
