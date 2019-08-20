package main

import (
	"citicup-admin/common"
	"citicup-admin/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"net/http"
	"strconv"
)

//将所有方法加入路由
func registerRouter(router *gin.Engine) {
	new(controller.UserController).Router(router)
}
//config配置读取
func initConfig(){
	cfg := new(common.Config)
	cfg.Parse("config/app.properties")
	fmt.Println("[ok] load config ")
	common.SetCfg(cfg)

	common.LoggerConfiguration(cfg.Logger["filepath"])

	gin.SetMode(cfg.App["mode"])

	for k, ds := range cfg.Datasource {
		e, err := xorm.NewEngine(ds["driveName"], ds["dataSourceName"])
		if err != nil {
			return
		}
		e.ShowSQL(ds["showSql"] == "true")

		n, _ := strconv.Atoi(ds["maxIdle"])
		e.SetMaxIdleConns(n)
		n, _ = strconv.Atoi(ds["maxOpen"])
		e.SetMaxOpenConns(n)

		common.SetEngine(k, e)
	}
	router := gin.Default()

	registerRouter(router)

	err := http.ListenAndServe(":"+cfg.App["port"], router)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("[ok] app run", cfg.App["addr"]+":"+cfg.App["port"])
	}
}

func main() {
	initConfig()

}
