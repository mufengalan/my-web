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
	if err := configor.Load(&conf, "/etc/config.yaml"); err != nil {
		panic(err)
	}
}
