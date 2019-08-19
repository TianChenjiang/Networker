package model

import (
	"context"
	"networker/backend/src/model/impl/entity"
	"networker/backend/src/pkg/database"
	"networker/backend/src/schema"
)

func NewUser(db *database.DB) *User {
	return &User{db}
}

type User struct {
	db *database.DB
}


func (a *User) Create(ctx context.Context, item schema.User) error {
	sitem := entity.SchemaUser(item)
	user := sitem.ToUser()
	a.db.Create(&user)

	return nil
}











