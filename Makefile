build:
	docker build -t golang-backend-image:latest .

clean:
	docker system prune --force
	docker image prune --force

init:
	docker build -t golang-backend-image .
	docker compose up

be-bash:
	docker exec -it manabie-be bash

run:
	docker compose up
run-no-logs:
	docker compose up -d