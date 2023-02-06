build:
	go build -o imagePreviewr cmd/app/main.go

run:
	docker-compose up

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

test:
	go test -v -race -count 1 ./...

test_api:
	docker-compose -f ./docker-compose.test.yml up -d
	go test -v --tags=integration ./integration/test/...
	docker-compose -f ./docker-compose.test.yml down