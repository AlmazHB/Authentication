package main

import (
	"fmt"
	"log"

	"github.com/diamondhulk625/web"
	"github.com/diamondhulk625/web/internall/app/configs"
	"github.com/diamondhulk625/web/internall/app/database"
	"github.com/diamondhulk625/web/internall/app/handler"
	"github.com/diamondhulk625/web/internall/app/repository"
	"github.com/diamondhulk625/web/internall/app/service"
	"github.com/spf13/viper"
)

const (
	confPath = "configs"
	confName = "config"
)

func main() {
	if err := configs.Init(confPath, confName); err != nil {
		log.Fatalf("Error instalizing confis:%s", err.Error())
	}

	if err := database.Init(viper.GetString("uri"), viper.GetString("dbName"), viper.GetString("dbCollection")); err != nil {
		log.Fatalf("Error connecting to MongDB:%s", err.Error())
	}

	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatalf("Error MongDB not Colsed!!!")
		}
	}()
	fmt.Printf("Server is listinig on %s port.\n", viper.GetString("port"))
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	server := new(web.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
