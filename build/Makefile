GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
GODEP=godep save
VERSION?=

APP_NAME?=
WORKSPACE?=$(GOPATH)/src/github.com/rosenlo/alertmanager-webhook-wechat
BIN_PATH?=$(WORKSPACE)/build/bin
APP_PATH=$(WORKSPACE)/cmd
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date "+%Y-%m-%dT%H:%M:%S%z")
LDFLAGS += -X $(shell $(GOLIST) $(APP_PATH)/$(APP_NAME)).GitHash=$(GIT_COMMIT)
LDFLAGS += -X $(shell $(GOLIST) $(APP_PATH)/$(APP_NAME)).Version=$(VERSION)
LDFLAGS += -X $(shell $(GOLIST) $(APP_PATH)/$(APP_NAME)).BuildTime=$(BUILD_TIME)
GOFILES = $(shell find $(WORKSPACE) -type f -name '*.go')
GOFILES_ALL = $(GOFILES)

${APP_NAME}: $(GOFILES_ALL)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v -ldflags "$(LDFLAGS)" -o $(BIN_PATH)/$(APP_NAME)
