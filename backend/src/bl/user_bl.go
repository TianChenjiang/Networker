package bl

import (
	"context"
	"networker/backend/src/schema"
)

type IUser interface  {
	// 创建数据
	Create(ctx context.Context, item schema.User) (*schema.User, error)
}