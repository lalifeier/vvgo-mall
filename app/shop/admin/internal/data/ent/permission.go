// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/permission"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID int64 `json:"id,omitempty"`
	// MenuID holds the value of the "menu_id" field.
	// 菜单ID
	MenuID int64 `json:"menu_id,omitempty"`
	// Name holds the value of the "name" field.
	// 权限名称
	Name string `json:"name,omitempty"`
	// Permission holds the value of the "permission" field.
	// 权限
	Permission string `json:"permission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID, permission.FieldMenuID:
			values[i] = new(sql.NullInt64)
		case permission.FieldName, permission.FieldPermission:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Permission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int64(value.Int64)
		case permission.FieldMenuID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				pe.MenuID = value.Int64
			}
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldPermission:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permission", values[i])
			} else if value.Valid {
				pe.Permission = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return (&PermissionClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", menu_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.MenuID))
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", permission=")
	builder.WriteString(pe.Permission)
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission

func (pe Permissions) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
