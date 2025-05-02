package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"user-service/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	//получаем путь к миграциям и строку подключения к БД
	migrationsPath := cfg.Migrations.Path
	dbUrl := cfg.PG.MigrationsUrl()

	//создаем объект миграции
	m, err := migrate.New(migrationsPath, dbUrl)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	//выполняем команду, переданную как аргумент (up, down)
	if len(os.Args) < 2 {
		log.Fatalf("Usage: migrate <command>\nAvailable commands: up, down")
	}

	command := os.Args[1]

	//выполняем миграции в зависимости от команды
	switch command {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("failed to apply migrations: %v", err)
		}
		fmt.Print("migrations applied successfully\n")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("failed to rollback migrations: %v", err)
		}
		fmt.Print("migrations rolled back successfully\n")
	default:
		log.Fatalf("unknown command: %s\n", command)
	}
}
