package main

import (
	"github.com/mufengalan/my-web/config/autoload"
	"github.com/mufengalan/my-web/internal/routers"
)

func main() {
	//config.Load()
	autoload.InitMysql()
	routers.InitRouter()
}
