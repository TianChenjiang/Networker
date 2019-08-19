package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"networker/backend/src/pkg/container"
	"networker/backend/src/pkg/setting"
	"networker/backend/src/pkg/util"
	"networker/backend/src/routers"
)

func init() {
	setting.Setup()
	util.Setup()
}


func main() {
	app := gin.New()
	gin.SetMode(setting.ServerSetting.RunMode)

	container := container.BuildContainer()
	fmt.Println(container)
	err := routers.RegisterRouter(app, container)
	if err != nil {
		fmt.Println(err)
	}

	routersInit := app
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20


	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}