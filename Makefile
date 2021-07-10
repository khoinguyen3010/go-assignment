build:
	docker build -t golang-backend-image .
run-logs:
	docker compose up
run-no-logs:
	docker compose up -d
sys-prune:
	docker system prune --force
img-prune:
	docker image prune --force