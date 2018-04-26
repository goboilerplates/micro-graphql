GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

all: install

install:
	go get github.com/goboilerplates/micro-rest
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	bash script/coverage.sh

build:
	bash script/BuildMulti.sh

buildDocker:
	docker build -t goboilerplates/micro-rest .