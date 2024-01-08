build:
	go mod tidy
	docker-compose up -d
build_with_logs:
	go mod tidy
	docker-compose up
down:
	docker-compose down