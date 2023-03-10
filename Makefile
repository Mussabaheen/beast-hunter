test:
	$(info ๐งช Testing...)
	go test ./...

lint:
	$(info ๐งช Checking the lint)
	golangci-lint run ./...

format:
	$(info ๐๏ธ formatting...)
	@go fmt ./...

build: clean download
	$(info ๐ฆ Building...)
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