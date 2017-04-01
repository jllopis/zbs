# Copyright 2017 Joan Llopis. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
.PHONY: help
.DEFAULT_GOAL := help

BLDDIR = _build
BINDIR = _bin
BLDDATE=$(shell date -u +%Y%m%dT%H%M%S)
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1` || echo v0.0.1)
REVISION=$(shell git rev-list --all --max-count=1 --abbrev-commit)
LDFLAGS=" -s -X main.BuildDate='${BLDDATE}' -X main.Version='${VERSION}' -X main.Revision='${REVISION}'"
SRCS = $(wildcard *.go ./**/*.go)

DOCKER=$(shell which docker)
COMPOSE=$(shell which docker-compose)

BINNAME="zbs"
PROJECT="acb-apis"
GITPROJECT="zbs"
ORG_PATH=github.com/jllopis
REPO_PATH=$(ORG_PATH)/$(GITPROJECT)

export PATH := $(PWD)/_bin:$(PATH)

dep: ## Vendor go dependencies
	@echo "Vendoring dependencies"
	@go get -u github.com/FiloSottile/gvt
	@gvt rebuild

$(BLDDIR):
	mkdir ${BLDDIR} || true

$(BINDIR):
	mkdir ${BINDIR} || true

linux: $(BLDDIR) ## build linux amd64 binary
	$(shell GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags ${LDFLAGS} -a -installsuffix cgo \
	-o ${BLDDIR}/${BINNAME}_${VERSION}_srv_linux_amd64.bin . \
	&& chmod +x ${BLDDIR}/${BINNAME}_${VERSION}_srv_linux_amd64.bin \
	)

osx: $(BLDDIR) ## build darwin amd64 binary
	$(shell GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -ldflags ${LDFLAGS} -a -installsuffix cgo \
	-o ${BLDDIR}/${BINNAME}_${VERSION}_srv_darwin_amd64.bin . \
	&& chmod +x ${BLDDIR}/${BINNAME}_${VERSION}_srv_darwin_amd64.bin \
	)

docker: linux ## build a docker image from the linux binary
	$(DOCKER) build --no-cache -t eu.gcr.io/${PROJECT}/${BINNAME}:${VERSION} .

push: docker ## push the generated docker image to google project ${PROJECT} repository
	gcloud --project=${PROJECT} docker -- push eu.gcr.io/${PROJECT}/${BINNAME}:${VERSION}

acb: docker ## build a docker image from the linux binary targetting the acb private repository
	$(DOCKER) tag eu.gcr.io/${PROJECT}/${BINNAME}:${VERSION} docker.acb.info:5000/${BINNAME}:${VERSION}

pushacb: acb ## push the generated acb docker image to private ACB repository
	$(DOCKER) push docker.acb.info:5000/${BINNAME}:${VERSION}

save-image: docker ## save the image of the project as a filesystem tar image to be loaded in docker
	$(DOCKER) -o ${BINNAME}:${VERSION}.tar save eu.gcr.io/${PROJECT}/${BINNAME}:${VERSION}

pb: $(BINDIR) _bin/protoc _bin/protoc-gen-go ## compile the protocol buffers files into resources
	_bin/protoc -I/usr/local/include -I. \
		-I$$PWD/vendor \
		-I$$PWD/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		services/*.proto ;

gw:  $(BINDIR) _bin/protoc _bin/protoc-gen-go _bin/protoc-gen-grpc-gateway ## build the REST gateway resource files
	@_bin/protoc -I/usr/local/include -I. \
		-I$$PWD/vendor \
		-I$$PWD/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		services/*.proto ;

doc: $(BINDIR) _bin/protoc _bin/protoc-gen-go _bin/protoc-gen-grpc-gateway _bin/protoc-gen-swagger ## build the documentation files: swagger json
	@_bin/protoc -I/usr/local/include -I. \
		-I$$PWD/vendor \
		-I$$PWD/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		services/*.proto ;

allpb: pb gw ## build proto and gateway files

_bin/protoc:
	@./scripts/get-protoc ${BINDIR}/protoc

_bin/protoc-gen-go:
	@go build -o ${BINDIR}/protoc-gen-go $(REPO_PATH)/vendor/github.com/golang/protobuf/protoc-gen-go

_bin/protoc-gen-grpc-gateway:
	@go build -o ${BINDIR}/protoc-gen-grpc-gateway $(REPO_PATH)/vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

_bin/protoc-gen-swagger:
	@go build -o ${BINDIR}/protoc-gen-swagger $(REPO_PATH)/vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

test: linux ## build the image and run. Uses file docker-compose.prod.yml
	@$(COMPOSE) -f docker-compose.prod.yml up

run-dev: ## force rebuild the docker-compose image (even if they haven't changed) and run using docker-compose.yml
	@docker-compose --build up -d

clean: ## remove the generated files to start clean but keep the images
	@rm -v ${BLDDIR}/${BINNAME}_srv_linux || echo "linux version does not exist"
	@rm -v ${BLDDIR}/${BINNAME}_srv_darwin || echo "darwin version does not exist"
	@rm -rf ${BLDDIR}/ || echo "_build directory do not exist"
	@rm -rf ${BINDIR}/ || echo "_bin directory do not exist"
	$(COMPOSE) stop
	$(COMPOSE) rm

clean-images: ## remove the generated images
	@$(DOCKER) rmi eu.gcr.io/${PROJECT}/${BINNAME}:${VERSION} || true
	@$(DOCKER) rmi docker.acb.info:5000/${BINNAME}:${VERSION} || true

clean-all: clean clean-images ## remove the generated files AND the images

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

