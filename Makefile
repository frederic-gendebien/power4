.PHONY: build

PROJECT=power4
VERSION=$$(git rev-parse --short=10 HEAD)

clean:
	go clean -cache

build:
	go build -v ./...

run:
	go run cmd/main.go

container:
	docker build -f build/Dockerfile . -t $(PROJECT):$(VERSION)

container-run: container
	docker run -it --rm $(PROJECT):$(VERSION)