run:
	go run main.go

swag:
	swag init

test:
	go test --cover ./...