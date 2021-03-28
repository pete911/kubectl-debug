test:
	GO111MODULE=on go test ./...
.PHONY:test

build: test
	GO111MODULE=on go build
.PHONY:build

install: test
	GO111MODULE=on go install
.PHONY:install
