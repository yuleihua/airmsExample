## commom Makefile

DATETIME=`date +%FT%T%z`
PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

all: fmt build

.PHONY: fmt vet build

list:
	@echo ${DATETIME}
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

proto:
	#go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . node/apis/airmsExample.proto --go_out=plugins=grpc:${GOPATH}/src

build: proto
	go build -o dist/airmsExample main.go

clean:
	@if [ -f dist/airmsExample ] ; then rm dist/airmsExample ; fi


