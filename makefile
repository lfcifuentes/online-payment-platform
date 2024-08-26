include .env

docker_up:
	docker-compose up --build -d

docker_down:
	docker-compose down
