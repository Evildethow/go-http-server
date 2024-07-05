default: build

.PHONY: build
build: fmt tidy test
	CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/server cmd/server/main.go

.PHONY: run
run: build
	bin/server

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf bin

.PHONY: release
release: build
	docker build --tag evildethow/httpserver .

.PHONY: docker-run
docker-run: release
	docker rm -f gohttpserver 2>/dev/null
	docker run -p 8080:8080 --name gohttpserver evildethow/httpserver