package main

import (
	"github.com/AlmazHB/Authentication/internall/app/configs"
	"github.com/AlmazHB/Authentication/internall/app/database"
	"github.com/AlmazHB/Authentication/internall/app/handler"
	"github.com/AlmazHB/Authentication/internall/app/repository"
	"github.com/AlmazHB/Authentication/internall/app/service"
	"github.com/AlmazHB/Authentication/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	confPath = "configs"
	confName = "config"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := configs.Init(confPath, confName); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := database.Init(viper.GetString("uri"), viper.GetString("dbName")); err != nil {
		logrus.Fatalf("Error connecting to MongoDB: %s", err.Error())
	}

	defer func() {
		err := database.Close()
		if err != nil {
			logrus.Fatalf("Error MongoDB not Closed!!!")
		}
	}()

	db := database.GetDatabase()
	logrus.Printf("Server is listening on %s port.\n", viper.GetString("port"))
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(web.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
