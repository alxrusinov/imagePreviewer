build:
	docker-compose build

run:
	docker-compose up

lint:
	golangci-lint run

test:
	go test -v -race -count 1 ./...

test_api:
	docker-compose -f ./docker-compose.test.yml up -d
	go test -v --tags=integration ./integration/test/...
	docker-compose -f ./docker-compose.test.yml down