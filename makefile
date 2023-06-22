GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=tree-visualizer
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_PATH= ./test/...
BUILD_PATH=./build
GOGENERATE=${GOCMD} generate

.PHONY: all test clear build

all: 
	test build
build:
	${GOGET}
	mkdir -p ${BUILD_PATH}
	${GOBUILD} -o ${BUILD_PATH}/${BINARY_NAME}

build-all:
	make build
	cd ./comunica-js/comunica-feature-link-traversal && yarn install
	cd ./comunica-js && yarn install

test: 
	cd test/communication && yarn install
	$(GOTEST) -v ${TEST_PATH}
clear: 
	rm -f ${BUILD_PATH}/$(BINARY_NAME)
	rm -f ${BUILD_PATH}/$(BINARY_UNIX)
run:
	${BUILD_PATH}/$(BINARY_NAME)
build-run:
	make build
	make run