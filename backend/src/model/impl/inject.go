package impl

import (
	"go.uber.org/dig"
	"networker/backend/src/model"
	"networker/backend/src/model/impl/entity"
	imodel "networker/backend/src/model/impl/model"
	"networker/backend/src/pkg/database"
)

// AutoMigrate 自动映射数据表
func AutoMigrate(db *database.DB) error {
	return db.AutoMigrate(

		new(entity.User),

	).Error
}


func Inject(container *dig.Container) error {

	container.Provide(imodel.NewUser, dig.As(new(models.IUser)))
	return nil
}
