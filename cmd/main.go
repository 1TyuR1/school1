package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Запуск сервера через Echo
	if err := e.Start(":9091"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}

	//handlers := new(handler.Handler)

	//srv := new(school1.Server)
	//if err := srv.Run("9091", handlers.InitRoutes()); err != nil {
	//	log.Fatalf("Ошибка запуска сервера: %s", err.Error())

	//}

	// echo/v4

}
