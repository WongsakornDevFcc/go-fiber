APP_NAME = apiserver
BUILD_DIR = $(PWD)/app
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost:5432/fiber-base?sslmode=disable

start:
	docker compose up -d

stop:
	docker compose down

rebuild:
	docker compose up -d --build

restart:
	make stop
	make start
	
remove-volume:
	docker-compose down -v

logs:
	docker-compose logs -f

migrate.create: 
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $(name) 
# make migrate.create name=....

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)


# migrate -path ./platform/migrations -database "postgres://postgres:password@localhost:5432/fiber-base?sslmode=disable" up