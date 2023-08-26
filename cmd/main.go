package main

import (
	"github/todo-app"
	"github/todo-app/pkg/handler"
	"github/todo-app/pkg/repository"
	"github/todo-app/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error during running http server: %s", err.Error())
	}
}
