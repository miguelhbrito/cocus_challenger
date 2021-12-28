ENV = $(shell go env GOPATH)
GO_VERSION = $(shell go version)
GO111MODULE=on
BINARY_NAME=cocus.out

# Look for versions prior to 1.10 which have a different fmt output
# and don't lint with gofmt against them.
ifneq (,$(findstring go version go1.8, $(GO_VERSION)))
	FMT=
else ifneq (,$(findstring go version go1.9, $(GO_VERSION)))
	FMT=
else
    FMT=--enable gofmt
endif

lint: # @HELP lint files and format if possible
	@echo "executing linter"
	gofmt -s -w .
	GO111MODULE=on golangci-lint run -c .golangci-lint.yml $(FMT) ./...
dep-linter: # @HELP install the linter dependency
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(ENV)/bin $(GOLANG_CI_VERSION)
build-simple:
	go build -o bin/${BINARY_NAME} cmd/main.go
build: # @HELP build the packages
	chmod +x scripts/build-simple.sh
	./scripts/build-simple.sh
run-cocus-gateway-build:
	go build -o bin/${BINARY_NAME} cmd/main.go
	./bin/${BINARY_NAME}
run-cocus-gateway:
	echo "running the api server"
	chmod +x scripts/run-server.sh
	./scripts/run-server.sh
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
ci: # @HELP executes on CI
ci: deps test fuzz dep-linter lint

gh-ci: # @HELP executes on GitHub Actions
gh-ci: deps test dep-linter lint

all: deps test fuzz lint