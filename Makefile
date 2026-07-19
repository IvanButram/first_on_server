include .env
export

run:
	go run main.go

compose:
	docker compose up -d --build

uncompose:
	docker compose down

compose_logs:
	docker compose logs -f