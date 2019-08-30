package main

import (
	"citicup-admin/internal/config"
	"citicup-admin/internal/service"
	"citicup-admin/internal/web"
	"context"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

)
// @title Citicup Api
// @version 1.0
// @description The Official Api Document For the Citicup

// @contact.name API Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	//初始化配置文件
	c := config.InitConfig()
	//建立容器
	srv := service.New(c)
	httpSrv := web.New(c, srv)
	//end
	log.Print("set up the container")
	cycle(httpSrv)

}

func cycle(httpSrv *http.Server) {
	t := make(chan os.Signal, 1)
	for {
		s := <-t
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Print("error")
			}
			httpSrv.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
