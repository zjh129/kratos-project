// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-project/app/grpc_user/internal/data/ent/userinfo"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserInfo is the model entity for the UserInfo schema.
type UserInfo struct {
	config `json:"-"`
	// ID of the ent.
	// 用户ID
	ID int `json:"id,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// 用户账号
	Account string `json:"account,omitempty"`
	// 用户密码
	Password string `json:"password,omitempty"`
	// 用户名称
	Name string `json:"name,omitempty"`
	// 用户头像地址
	Avatar string `json:"avatar,omitempty"`
	// 用户类别
	Type int32 `json:"type,omitempty"`
	// 可用状态(0:禁用,1:启用)
	StatusIs int32 `json:"status_is,omitempty"`
	// 添加时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userinfo.FieldID, userinfo.FieldType, userinfo.FieldStatusIs:
			values[i] = new(sql.NullInt64)
		case userinfo.FieldAccount, userinfo.FieldPassword, userinfo.FieldName, userinfo.FieldAvatar:
			values[i] = new(sql.NullString)
		case userinfo.FieldDeleteTime, userinfo.FieldCreatedAt, userinfo.FieldUpdatedAt, userinfo.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserInfo fields.
func (ui *UserInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ui.ID = int(value.Int64)
		case userinfo.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				ui.DeleteTime = value.Time
			}
		case userinfo.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				ui.Account = value.String
			}
		case userinfo.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				ui.Password = value.String
			}
		case userinfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ui.Name = value.String
			}
		case userinfo.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				ui.Avatar = value.String
			}
		case userinfo.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ui.Type = int32(value.Int64)
			}
		case userinfo.FieldStatusIs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status_is", values[i])
			} else if value.Valid {
				ui.StatusIs = int32(value.Int64)
			}
		case userinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ui.CreatedAt = value.Time
			}
		case userinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ui.UpdatedAt = value.Time
			}
		case userinfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ui.DeletedAt = value.Time
			}
		default:
			ui.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserInfo.
// This includes values selected through modifiers, order, etc.
func (ui *UserInfo) Value(name string) (ent.Value, error) {
	return ui.selectValues.Get(name)
}

// Update returns a builder for updating this UserInfo.
// Note that you need to call UserInfo.Unwrap() before calling this method if this UserInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ui *UserInfo) Update() *UserInfoUpdateOne {
	return NewUserInfoClient(ui.config).UpdateOne(ui)
}

// Unwrap unwraps the UserInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ui *UserInfo) Unwrap() *UserInfo {
	_tx, ok := ui.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserInfo is not a transactional entity")
	}
	ui.config.driver = _tx.drv
	return ui
}

// String implements the fmt.Stringer.
func (ui *UserInfo) String() string {
	var builder strings.Builder
	builder.WriteString("UserInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ui.ID))
	builder.WriteString("delete_time=")
	builder.WriteString(ui.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("account=")
	builder.WriteString(ui.Account)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(ui.Password)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ui.Name)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(ui.Avatar)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ui.Type))
	builder.WriteString(", ")
	builder.WriteString("status_is=")
	builder.WriteString(fmt.Sprintf("%v", ui.StatusIs))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ui.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ui.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ui.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserInfos is a parsable slice of UserInfo.
type UserInfos []*UserInfo
