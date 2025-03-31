package main

import (
	"github.com/mufengalan/my-web/config"
	"github.com/mufengalan/my-web/config/autoload"
	"github.com/mufengalan/my-web/internal/routers"
)

func main() {
	logger := autoload.InitZapLogger()
	defer logger.Sync() // 确保退出前刷新日志
	config.Load()
	cnf := autoload.MysqlConfig{
		Host:     config.Items().Mysql.Host,
		Port:     config.Items().Mysql.Port,
		Username: config.Items().Mysql.Username,
		Password: config.Items().Mysql.Password,
	}
	autoload.InitMysql(cnf)
	routers.InitRouter(logger)
}
