// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/dictdata"
)

// DictData is the model entity for the DictData schema.
type DictData struct {
	config `json:"-"`
	// ID of the ent.
	// 字典数据id
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	// 更新人
	CreatedBy int64 `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	// 更新人
	UpdatedBy int64 `json:"updated_by,omitempty"`
	// DictTypeID holds the value of the "dict_type_id" field.
	// 字典类型id
	DictTypeID int64 `json:"dict_type_id,omitempty"`
	// Label holds the value of the "label" field.
	// 字典标签
	Label string `json:"label,omitempty"`
	// Value holds the value of the "value" field.
	// 字典键值
	Value string `json:"value,omitempty"`
	// Sort holds the value of the "sort" field.
	// 排序
	Sort int8 `json:"sort,omitempty"`
	// Status holds the value of the "status" field.
	// 状态 0:禁用 1:启用 -1:删除
	Status int8 `json:"status,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	// 是否默认值 0:否 1:是
	IsDefault int8 `json:"is_default,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DictData) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictdata.FieldID, dictdata.FieldCreatedBy, dictdata.FieldUpdatedBy, dictdata.FieldDictTypeID, dictdata.FieldSort, dictdata.FieldStatus, dictdata.FieldIsDefault:
			values[i] = new(sql.NullInt64)
		case dictdata.FieldLabel, dictdata.FieldValue, dictdata.FieldRemark:
			values[i] = new(sql.NullString)
		case dictdata.FieldCreatedAt, dictdata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DictData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DictData fields.
func (dd *DictData) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dd.ID = int64(value.Int64)
		case dictdata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dd.CreatedAt = value.Time
			}
		case dictdata.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				dd.CreatedBy = value.Int64
			}
		case dictdata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dd.UpdatedAt = value.Time
			}
		case dictdata.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				dd.UpdatedBy = value.Int64
			}
		case dictdata.FieldDictTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dict_type_id", values[i])
			} else if value.Valid {
				dd.DictTypeID = value.Int64
			}
		case dictdata.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				dd.Label = value.String
			}
		case dictdata.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				dd.Value = value.String
			}
		case dictdata.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				dd.Sort = int8(value.Int64)
			}
		case dictdata.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dd.Status = int8(value.Int64)
			}
		case dictdata.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				dd.Remark = value.String
			}
		case dictdata.FieldIsDefault:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				dd.IsDefault = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DictData.
// Note that you need to call DictData.Unwrap() before calling this method if this DictData
// was returned from a transaction, and the transaction was committed or rolled back.
func (dd *DictData) Update() *DictDataUpdateOne {
	return (&DictDataClient{config: dd.config}).UpdateOne(dd)
}

// Unwrap unwraps the DictData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dd *DictData) Unwrap() *DictData {
	tx, ok := dd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DictData is not a transactional entity")
	}
	dd.config.driver = tx.drv
	return dd
}

// String implements the fmt.Stringer.
func (dd *DictData) String() string {
	var builder strings.Builder
	builder.WriteString("DictData(")
	builder.WriteString(fmt.Sprintf("id=%v", dd.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(dd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", dd.CreatedBy))
	builder.WriteString(", updated_at=")
	builder.WriteString(dd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", dd.UpdatedBy))
	builder.WriteString(", dict_type_id=")
	builder.WriteString(fmt.Sprintf("%v", dd.DictTypeID))
	builder.WriteString(", label=")
	builder.WriteString(dd.Label)
	builder.WriteString(", value=")
	builder.WriteString(dd.Value)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", dd.Sort))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", dd.Status))
	builder.WriteString(", remark=")
	builder.WriteString(dd.Remark)
	builder.WriteString(", is_default=")
	builder.WriteString(fmt.Sprintf("%v", dd.IsDefault))
	builder.WriteByte(')')
	return builder.String()
}

// DictDataSlice is a parsable slice of DictData.
type DictDataSlice []*DictData

func (dd DictDataSlice) config(cfg config) {
	for _i := range dd {
		dd[_i].config = cfg
	}
}
