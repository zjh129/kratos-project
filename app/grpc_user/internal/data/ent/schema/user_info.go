// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserInfo struct {
	ent.Schema
}

func (UserInfo) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Comment("用户ID"), field.String("account").Comment("用户账号"), field.String("password").Comment("用户密码"), field.String("name").Comment("用户名称"), field.String("avatar").Comment("用户头像地址"), field.Int32("type").Comment("用户类别"), field.Int32("status_is").Comment("可用状态(0:禁用,1:启用)"), field.Time("created_at").Comment("添加时间"), field.Time("updated_at").Comment("修改时间"), field.Time("deleted_at").Optional().Comment("删除时间")}
}
func (UserInfo) Edges() []ent.Edge {
	return nil
}
func (UserInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "user_info"}}
}

// Mixin of the UserInfo.
func (UserInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}
