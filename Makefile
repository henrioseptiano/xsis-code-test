build:
	go mod tidy
	docker-compose up -d
build_with_logs:
	go mod tidy
	docker-compose up
down:
	docker-compose down
test:
	go test -v ./...
test_cover:
	go test -coverprofile=coverage.out ./app/handlers ./app/repository ./app/usecase ./utils
	go tool cover -func=coverage.out
test_cover_html:
	go test -coverprofile=coverage.out ./app/handlers ./app/repository ./app/usecase ./utils
	go tool cover -html=coverage.out