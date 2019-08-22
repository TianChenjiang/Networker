package impl

import (
	"go.uber.org/dig"
	"networker/backend/src/bl"
	"networker/backend/src/bl/impl/internal"
)


func Inject(container *dig.Container) error {

	container.Provide(internal.NewUser, dig.As(new(bl.IUser)))
	return nil
}