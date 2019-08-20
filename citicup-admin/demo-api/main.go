package main

import (
	"citicup-admin/demo-api/internal/conf"
	"citicup-admin/demo-api/internal/server/http"
	"citicup-admin/demo-api/internal/service"
	"context"
	"flag"
	"github.com/CaribouW/gin-common-lib/library/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()

	// IDE中，你也可以放开注释，直接运行。
	/*dir, _ := filepath.Abs("./app/demo-api/configs/application.toml")
	flag.Set("conf", dir)*/

	// 初始化配置
	conf.Init()
	// 初始化日志
	log.Init(conf.Conf.Log)
	defer log.Close()
	srv := service.New(conf.Conf)
	log.Info("github.com/CaribouW/gin-common-lib start")
	httpSrv := http.New(conf.Conf, srv)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("github.com/CaribouW/gin-common-lib exit")
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
