build:
	docker-compose build app
run:
	docker-compose up app
down:
	docker-compose down
swag:
	swag init -g cmd/main.go -o docs
linter:
	golangci-lint run