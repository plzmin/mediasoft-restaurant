build:
	docker compose up --build

docker-down:
	docker compose down --volumes

migrate-up:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/customer?sslmode=disable up


migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/customer?sslmode=disable down
