BINARY_NAME=webapp.exe

all: build run

build:
	@echo "building application..."
	@go build -o webapp.exe ./cmd/web
	@echo "Webapp built"

run:
	@echo "running webapp..."
	@./${BINARY_NAME}

clean:
	@echo "cleaning..."
	@go clean
	@if exist ${BINARY_NAME} del ${BINARY_NAME}
	@echo "cleaned"

test:
	@echo "testing..."
	@go test ./..
	@echo "DONE"
