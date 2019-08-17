package app

import (
	"context"
)

type options struct {
	ConfigFile string
	ModelFile  string
	WWWDir     string
	SwaggerDir string
	Version    string
}

type Option func(*options)

func Setup(ctx context.Context, opts ...Option)  {


}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

//func BuildContainer() (*dig.Container, func()) {
//	// 创建依赖注入容器
//	container := dig.New()
//
//	// 注入casbin
//	//container.Provide(NewEnforcer)
//
//	// 注入认证模块
//	//auther, err := InitAuth()
//	//handleError(err)
//	//container.Provide(func() auth.Auther {
//	//	return auther
//	//})
//
//
//
//	// 注入bll
//	err = impl.Inject(container)
//	handleError(err)
//
//	return container, func() {
//		if auther != nil {
//			auther.Release()
//		}
//		if storeCall != nil {
//			storeCall()
//		}
//	}
//}