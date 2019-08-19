package data

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	inj "networker/backend/src/model/impl"
	"networker/backend/src/pkg/database"
)

func InitStore(container *dig.Container) error {
	db := database.Setup()
	container.Provide(func() *gorm.DB{
		return db
	})

	inj.Inject(container)

	return nil
}
