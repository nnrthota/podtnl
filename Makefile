GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go mod download
	go build

.PHONY: test
test:
	 go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t podtnl:latest