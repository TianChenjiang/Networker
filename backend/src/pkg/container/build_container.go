package container

import (
	"go.uber.org/dig"
	"networker/backend/src/bl/impl"
	"networker/backend/src/pkg/data"
)

func handleError(err error) {
	if err != nil {
	panic(err)
	}
}

// BuildContainer 创建依赖注入容器
func BuildContainer() (*dig.Container){
	// 创建依赖注入容器
	container := dig.New()
	var err error

	// 注入casbin
	//container.Provide(NewEnforcer)

	// 注入认证模块
	//auther, err := InitAuth()
	//handleError(err)
	//container.Provide(func() auth.Auther {
	//	return auther
	//})

	// 注入存储模块
	err = data.InitStore(container)
	handleError(err)

	// 注入bl
	err = impl.Inject(container)
	handleError(err)

	return container
}


