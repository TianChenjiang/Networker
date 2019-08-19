package controller

import (
	"go.uber.org/dig"
)


func Inject(container *dig.Container) error {

	container.Provide(NewUser)
	return nil
}
