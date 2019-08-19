package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	ctl "networker/backend/src/routers/api/controller"
)

func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	fmt.Println(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(
		user *ctl.User,
	) error {

		g := app.Group("/api/v1")


		{
			g.POST("/users", user.CreateUser) //TODO
		}


		return nil
	})
}
