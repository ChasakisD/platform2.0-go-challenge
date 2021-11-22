update-docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	/Users/$(USER)/go/bin/swag init

clean:
	go clean

api-build:
	go build -o gwi/assignment

api-run:
	./gwi/assignment

api-test:
	go test -v ./...