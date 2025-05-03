ENTRY_POINT=cmd/app
MIGRATE_POINT=cmd/migrate
MIGRATIONS_DIR=migrations

# Генерация контрактов
proto-generate:
	protoc --proto_path=proto \
		--go_out=gen/user --go_opt=paths=source_relative \
		--go-grpc_out=gen/user --go-grpc_opt=paths=source_relative \
		proto/user.proto

# Тестирование
test:
	go test -v ./...

# Запуск
run:
	go run $(ENTRY_POINT)/main.go

# Запуск в тестовом режиме
run-dev:
	go run $(ENTRY_POINT)/main.go --dev

# Создание миграций
migrations-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Применение миграций
migrate-up:
	go run $(MIGRATE_POINT)/migrate.go up

# Откат миграций
migrate-down:
	go run $(MIGRATE_POINT)/migrate.go down