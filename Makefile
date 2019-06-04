PACKAGES := $(shell go list ./...)

export GO111MODULE=on

BUILD_TAGS := netgo
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w

BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test

build: dep_verify
	go build -mod=readonly ${BUILD_FLAGS} -o bin/vpn-node main.go

install: build
	mv bin/vpn-node ${GOPATH}/bin/

test:
	@go test -mod=readonly -cover ${PACKAGES}

benchmark:
	@go test -mod=readonly -bench=. ${PACKAGES}

dep_verify:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all build install test benchmark dep_verify