package controller

import (
	"go.uber.org/dig"
)

// Inject 注入ctl
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo *controller.Demo) {
//   })
func Inject(container *dig.Container) error {

	container.Provide(NewUser)
	return nil
}
