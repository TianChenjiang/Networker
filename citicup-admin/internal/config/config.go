package config

import (
	"flag"
	"fmt"
	config "github.com/go-ozzo/ozzo-config"
	"strings"
)

//配置文件相关
type Config struct {
	//App
	App *App
	//DB
	DataBase *DataBase
}

//App基本配置
type App struct {
	Port    int    //端口
	RunMode string //运行模式
}

//数据库配置
type DataBase struct {
	Type     string //类型
	User     string //用户名,一般为root
	Password string //密码
	Host     string //IP:PORT
	dbName   string //数据库名称
}

//解析配置文件内容
func initConfig(c *config.Config) *Config {
	var (
		appPrefix = "server"
		//dataPrefix = "data"
	)

	return &Config{
		App: &App{
			Port:    c.GetInt(appPrefix + ".port"),
			RunMode: c.GetString(appPrefix + ".runmode"),
		},
		DataBase: &DataBase{
			Type:     c.GetString("data.type"),
			User:     c.GetString("data.user"),
			Password: c.GetString("data.password"),
			Host:     c.GetString("data.host"),
			dbName:   c.GetString("data.dbname"),
		},
	}
}

func InitConfig() *Config {
	//默认为dev环境
	var profile = flag.String("profile", "dev", "profile环境,决定运行环境")
	flag.Parse()
	fmt.Print(*profile)
	c := config.New()
	//根据参数profile动态进行配置文件读取
	err := c.Load(strings.Join([]string{"configs/application-",
		*profile,
		".yml",
	}, ""))
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return initConfig(c)
}
