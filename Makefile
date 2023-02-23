test:
	$(info 🧪 Testing...)
	go test ./...

lint:
	$(info 🧪 Checking the lint)
	golangci-lint run ./...

format:
	$(info 🖊️ formatting...)
	@go fmt ./...

build: clean download
	$(info 📦 Building...)
	go build -o build/ ./...

clean:
	$(info Cleaning...)
	rm -rf ./build/* 

run: download
	$(info Running beast hunter...)
	go run cmd/beast-hunter/main.go

download:
	$(info Downloading go modules...)
	go mod tidy

start-docker:
	$(info Creating Docker Image)
	docker build -t beast-hunter .
	docker run beast-hunter