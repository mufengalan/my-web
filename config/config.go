package config

import (
	"github.com/jinzhu/configor"
	"github.com/mufengalan/my-web/config/autoload"
)

var Items = func() Conf {
	return conf
}
var conf Conf

type Conf struct {
	App   autoload.AppConfig
	Mysql autoload.MysqlConfig
}

func Load() {
	config := &configor.Config{ENVPrefix: "-"}
	if err := configor.New(config).Load(&conf, "/competition-tool-api/etc/config.yaml", "config.yaml", "../config.yaml"); err != nil {
		panic(err)
	}
}
