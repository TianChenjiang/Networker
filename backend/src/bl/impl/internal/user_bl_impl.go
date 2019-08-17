package internal

import (
	"context"
	models "networker/backend/src/model"
	"networker/backend/src/schema"
)

type User struct {
	UserModel models.IUser
}

func NewUser(user models.IUser,
	) *User {
	return &User{
		UserModel: user,
	}
}


// Create 创建数据
func (a *User) Create(ctx context.Context, item schema.User) (*schema.User, error) {

	err := a.checkUserName(ctx, item.Username)
	if err != nil {
		return nil, err
	}

	err = a.UserModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return nil, err
}



func (a *User) checkUserName(ctx context.Context, userName string) error{
	return nil
}







