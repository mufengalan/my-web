package main

import (
	"fmt"
	"github.com/mufengalan/my-web/config"
	"github.com/mufengalan/my-web/config/autoload"
	"github.com/mufengalan/my-web/internal/routers"
)

func main() {
	config.Load()
	cnf := autoload.MysqlConfig{
		Host:     config.Items().Mysql.Host,
		Port:     config.Items().Mysql.Port,
		Username: config.Items().Mysql.Username,
		Password: config.Items().Mysql.Password,
	}
	fmt.Println(cnf)
	autoload.InitMysql(cnf)
	routers.InitRouter()
}
