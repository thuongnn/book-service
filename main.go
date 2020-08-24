package main

import (
	"book-service/config"
	"book-service/src/api"
	"book-service/src/dao"
	"fmt"
	"github.com/astaxie/beego"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
)

func main() {
	config.Init()

	database, err := config.Database()
	if err != nil {
		log.Fatalf("failed to get database configuration: %v", err)
	}
	if err := dao.InitDatabase(database); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	registerRoutes()
	registerServiceWithConsul()
	beego.Run(config.GetAppPort())
}

func registerRoutes() {
	beego.Router("/healthcheck", &api.HomeAPI{}, "get:Get")
	beego.Router("/:id([0-9]+", &api.BookAPI{}, "get:Get;delete:Delete;put:Put")
	beego.Router("/all", &api.BookAPI{}, "get:List;post:Post")
}

func registerServiceWithConsul() {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "book-service"
	registration.Name = "book-service"
	address := hostname()
	registration.Address = address

	p := config.GetBookServicePort()
	registration.Port = p
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, p)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
