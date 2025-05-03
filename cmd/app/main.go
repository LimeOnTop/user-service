package main

import (
	"flag"
	"log"
	"user-service/config"
	"user-service/internal/app"
)

func main() {
	// Определение флага
	devMode := flag.Bool("dev", false, "Run server in development mode")
	flag.Parse()

	// Загрузка конфигурации
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	//	Запуск приложения
	app.Run(cfg, *devMode)
}
