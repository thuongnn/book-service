package main

import (
	"book-service/config"
	"book-service/src/api"
	"book-service/src/dao"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config.Init()

	database, err := config.Database()
	if err != nil {
		fmt.Errorf("failed to get database configuration: %v", err)
	}
	if err := dao.InitDatabase(database); err != nil {
		fmt.Errorf("failed to initialize database: %v", err)
	}

	registerRoutes()
	beego.Run(config.GetAppPort())
}

func registerRoutes() {
	beego.Router("/", &api.HomeAPI{}, "get:Get")
	beego.Router("/:id([0-9]+", &api.BookAPI{}, "get:Get;delete:Delete;put:Put")
	beego.Router("/", &api.BookAPI{}, "get:List;post:Post")
}
