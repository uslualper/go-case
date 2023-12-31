build:
	go build -o cmd/main main.go

run:
	go run main.go

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-restart:
	docker-compose restart

docker-exec:  
	docker exec -it ${CONTAINER_NAME} bash

swagger:
	swag init -g ./main.go ./docs

test:
	go test ./tests/...