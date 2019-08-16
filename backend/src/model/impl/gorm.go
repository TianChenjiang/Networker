package impl

import (
	"go.uber.org/dig"
	"networker/backend/src/model"
	"networker/backend/src/model/impl/entity"
	"networker/backend/src/model/impl/model"
	"networker/backend/src/pkg/gormplus"
)

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gormplus.DB) error {
	return db.AutoMigrate(

		new(entity.User),

	).Error
}

// Inject 注入gorm实现
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo IDemo) {
//   })
func Inject(container *dig.Container) error {

	container.Provide(model.NewUser, dig.As(new(models.IUser)))
	return nil
}
