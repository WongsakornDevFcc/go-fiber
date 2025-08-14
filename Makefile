build:
	docker-compose up --build

rebuild:
	docker-compose down & docker-compose up --build

logs:
	docker-compose logs -f

migrate:
	docker-compose exec api go run ./app/database/migrate.go

migrate-drop:
	docker-compose exec api go run ./app/database/migrate.go drop
