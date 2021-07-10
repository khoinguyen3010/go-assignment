build:
	docker build -t golang-backend-image .

prune-all:
	docker system prune --force
	docker image prune --force

init-proj:
	docker build -t golang-backend-image .
	docker compose up

be-bash:
	docker exec -it manabie-be bash

run-logs:
	docker compose up
run-no-logs:
	docker compose up -d