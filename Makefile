PKG_PREFIX := github.com/VictoriaMetrics/cloud-foundry-service-broker

DATEINFO_TAG ?= $(shell date -u +'%Y%m%d-%H%M%S')
BUILDINFO_TAG ?= $(shell echo $$(git describe --long --all | tr '/' '-')$$( \
	      git diff-index --quiet HEAD -- || echo '-dirty-'$$(git diff-index -u HEAD | openssl sha1 | cut -c 10-17)))

PKG_TAG ?= $(shell git tag -l --points-at HEAD)

GO_BUILDINFO = -X '$(PKG_PREFIX)/lib/buildinfo.Version=$(APP_NAME)-$(DATEINFO_TAG)-$(BUILDINFO_TAG)'

fmt:
	GO111MODULE=on gofmt -l -w -s ./lib
	GO111MODULE=on gofmt -l -w -s ./app

vet:
	GO111MODULE=on go vet ./lib/...
	GO111MODULE=on go vet ./app/...

golangci-lint: install-golangci-lint
	golangci-lint run --exclude '(SA4003|SA1019|SA5011):' -D errcheck -D structcheck --timeout 2m

install-golangci-lint:
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.46.2

errcheck: install-errcheck
	errcheck ./lib/...
	errcheck ./app/...

check-all: fmt vet lint errcheck golangci-lint

install-errcheck:
	which errcheck || GO111MODULE=off go get github.com/kisielk/errcheck

lint: install-golint errcheck
	golint lib/...
	golint app/...

install-golint:
	which golint || GO111MODULE=off go get golang.org/x/lint/golint

build:
	go build -ldflags "$(GO_BUILDINFO)" -o bin/cf-service-broker app/web/main.go