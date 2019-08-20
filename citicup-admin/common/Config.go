package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	App        map[string]string
	Datasource map[string](map[string]string)
	Logger     map[string]string
	All        map[string]string
}

func (cfg *Config) Parse(fpath string) {
	//初始化
	cfg.App = make(map[string]string)
	cfg.Datasource = make(map[string](map[string]string))
	cfg.All = make(map[string]string)
	cfg.Logger = make(map[string]string)

	fi, err := os.Open(fpath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		tmp := strings.TrimLeft(string(a), " ")
		tmp = strings.TrimRight(tmp, " ")
		if len(tmp) == 0 || strings.Index(tmp, "#") == 0 {
			continue
		}

		o := strings.Split(tmp, "=")
		if len(o) == 2 {
			cfg.All[o[0]] = o[1]
		} else if len(o) < 2 {
			cfg.All[o[0]] = ""
		} else {
			cfg.All[o[0]] = strings.TrimPrefix(tmp, o[0]+"=")

		}
	}

	//遍历
	for k, v := range cfg.All {
		if strings.Index(k, "restful.app.") == 0 {
			tmp := strings.TrimPrefix(k, "restful.app.")
			cfg.App[tmp] = v
		} else if strings.Index(k, "restful.logger.") == 0 {
			tmp := strings.TrimPrefix(k, "restful.logger.")
			cfg.Logger[tmp] = v
		} else if strings.Index(k, "restful.datasource.") == 0 {
			var sd = strings.Split(k, ".")
			if nil == cfg.Datasource[sd[2]] {
				cfg.Datasource[sd[2]] = make(map[string]string)
			}
			cfg.Datasource[sd[2]][sd[3]] = v

		}
	}

}

//获取整数,
func (cfg *Config) LoadCfg(key string) string {
	return cfg.All[key]
}

//获取字符串配置
func (cfg *Config) LoadString(key string) string {
	return cfg.All[key]
}

//获取整数,
func (cfg *Config) LoadInt(key string) (int, error) {
	return strconv.Atoi(cfg.All[key])
}

//获取32位整数
func (cfg *Config) LoadInt64(key string) (int64, error) {
	return strconv.ParseInt(cfg.All[key], 10, 64)
}

//获取64位整数
func (cfg *Config) LoadInt32(key string) (int64, error) {
	return strconv.ParseInt(cfg.All[key], 10, 32)
}

//获取布尔配置
func (cfg *Config) LoadBool(key string) bool {
	return cfg.All[key] == "true" || "TRUE" == cfg.All[key]
}

var _cfg *Config = nil

func SetCfg(c *Config) {
	_cfg = c
}
func GetCfg() *Config {
	return _cfg
}
