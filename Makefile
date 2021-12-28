ENV = $(shell go env GOPATH)
GO_VERSION = $(shell go version)
GO111MODULE=on
BINARY_NAME=cocus.out

build-simple:
	go build -o bin/${BINARY_NAME} app/cocus/main.go
build: # @HELP build the packages
	chmod +x platform/scripts/build-simple.sh
	./platform/scripts/build-simple.sh
run-cocus-gateway-build:
	go build -o bin/${BINARY_NAME} app/cocus/main.go
	./bin/${BINARY_NAME}
run-cocus-gateway:
	echo "running the api server"
	chmod +x platform/scripts/run-server.sh
	./platform/scripts/run-server.sh
config-up:
	docker-compose up -d
config-down:
	docker-compose down
clean:
	go clean
	rm bin/${BINARY_NAME}
test:
	go test -v ./...
test-cover:
	go test -cover ./...