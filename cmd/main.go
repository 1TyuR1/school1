package main

import (
	"log"

	"github.com/Fr0olyy/school1"
	"github.com/Fr0olyy/school1/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(school1.Server)
	if err := srv.Run("9091", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())

	}

}
