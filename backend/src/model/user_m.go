package models

import (
	"context"
	"networker/backend/src/schema"
)

//User 用户对象存储接口
type IUser interface {
	//根据ID查询
	Get(ctx context.Context, ID uint) (*schema.User, error)
	//创建用户数据
	Create(ctx context.Context, item schema.User) error
	//更新用户数据
	Update(ctx context.Context, ID uint, item schema.User) error
	//删除用户数据
	Delete(ctx context.Context, ID uint) error
	//修改密码
	UpdatePassword(ctx context.Context, ID uint, password string) error
}
