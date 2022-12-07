package main

import (
	"fmt"
	todo "github.com/NurzhanIT/Go_todo_api.git"
	"github.com/NurzhanIT/Go_todo_api.git/pkg/handler"
	"github.com/NurzhanIT/Go_todo_api.git/pkg/repository"
	"github.com/NurzhanIT/Go_todo_api.git/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error in initializing configs %s", err.Error())
	}
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server:\n %s", err.Error())
	} else {
		fmt.Println("started")
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
