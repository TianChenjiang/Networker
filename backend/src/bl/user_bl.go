package bl

import (
	"networker/backend/src/schema"
	"context"
)

type IUser interface  {
	// 创建数据
	Create(ctx context.Context, item schema.User) (*schema.User, error)
}