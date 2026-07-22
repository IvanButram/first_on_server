include .env
export

run:
	/app/exe

compose:
	docker compose up -d --build

uncompose:
	docker compose down

compose_logs:
	docker compose logs -f