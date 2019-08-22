package data

import (
	"go.uber.org/dig"
	inj "networker/backend/src/model/impl"
	"networker/backend/src/pkg/database"
)

func InitStore(container *dig.Container) error {
	db := database.Setup()
	container.Provide(func() *database.DB{
		return db
	})

	inj.Inject(container)

	return nil
}
