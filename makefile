GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=tree-visualiser
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_PATH= ./test/...
BUILD_PATH=./build
GOGENERATE=${GOCMD} generate

.PHONY: all test clear

all: 
	test build
build:
	${GOGET}
	git submodule update --init --recursive
	mkdir -p ${BUILD_PATH}
	${GOBUILD} -o ${BUILD_PATH}/${BINARY_NAME}
test: 
	$(GOTEST) -v ${TEST_PATH}
clear: 
	rm -f ${BUILD_PATH}/$(BINARY_NAME)
	rm -f ${BUILD_PATH}/$(BINARY_UNIX)
run:
	git submodule update --init --recursive
	mkdir -p ${BUILD_PATH}
	build
	${BUILD_PATH}/$(BINARY_NAME)