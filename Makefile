ENV = $(shell go env GOPATH)
GO_VERSION = $(shell go version)
GO111MODULE=on
BINARY_NAME=cocus.out

run-cocus-gateway-build:
	go build -o bin/${BINARY_NAME} app/cocus/main.go
	./bin/${BINARY_NAME}
run-cocus-gateway:
	echo "running the api server"
	chmod +x platform/scripts/run-server.sh
	./platform/scripts/run-server.sh
config-up:
	echo "starting up configs"
	docker-compose up -d
generate-secrets:
	echo "starting up secrets"
	chmod +x platform/scripts/vault/postsecret.sh
	chmod +x platform/scripts/vault/getsecret.sh
	./platform/scripts/vault/postsecret.sh
	./platform/scripts/vault/getsecret.sh
config-down:
	docker-compose down
clean:
	go clean
	rm bin/${BINARY_NAME}
test:
	go test -v ./...
test-cover:
	go test -cover ./...