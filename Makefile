.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/run-customer-service .

serve: build
	./bin/run-customer-service serve grpc

clean:
	rm -f ./bin/run-customer-service

dev:
	GOOS=linux make build
	docker build -t localhost:5000/run-customer-service .
	docker push localhost:5000/run-customer-service