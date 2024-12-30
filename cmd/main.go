package main

import (
	"github.com/mufengalan/my-web/config"
	"github.com/mufengalan/my-web/internal/routers"
)

func main() {
	config.Load()
	routers.InitRouter()
}
