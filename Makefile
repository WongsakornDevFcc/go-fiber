APP_NAME = apiserver
BUILD_DIR = $(PWD)/app
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost:5432/postgres?sslmode=disable

build:
	docker-compose up --build

rebuild:
	docker-compose down & docker-compose up --build

logs:
	docker-compose logs -f

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)
