package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/CaribouW/gin-common-lib/library/cache/redis"
	"github.com/CaribouW/gin-common-lib/library/database/sql"
	xhttp "github.com/CaribouW/gin-common-lib/library/http"
	"github.com/CaribouW/gin-common-lib/library/log"
	xtime "github.com/CaribouW/gin-common-lib/library/time"
)

var (
	confPath string
	Conf     = &Config{}
)

type Config struct {
	// App
	App *App
	// Log
	Log *log.Config
	// DB
	MySQL *MySQL
	// Redis
	Redis *redis.Config
	// HttpClient
	HttpClient *HttpClient
}

type App struct {
	HttpPort     int
	RunMode      string
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
}

type MySQL struct {
	School *sql.Config
}

type HttpClient struct {
	Paopao *xhttp.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
