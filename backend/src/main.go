package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"networker/backend/src/pkg/gormplus"
	"networker/backend/src/pkg/setting"
	"networker/backend/src/pkg/util"
	"networker/backend/src/routers"
)

func init() {
	setting.Setup()
	util.Setup()
	gormplus.Setup()
	//todo
}


func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	app := gin.New()
	err := routers.RegisterRouter(app, container)

	server := &http.Server{
		Addr:           endPoint,
		//Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}