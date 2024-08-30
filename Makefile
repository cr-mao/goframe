VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`
LDFLAGS = "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILDTIME)"


## tidy: 整理现有的依赖
.PHONY: tidy
tidy:
	go mod tidy

## download: go依赖下载
.PHONY: download
download:
	go mod download

## test: 单元测试全部测试代码
.PHONY: test
test:
	go test ./... -cover

## vet: 静态检测全部go代码
.PHONY: vet
vet:
#	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
#	go vet -vettool=`which shadow` ./...
	go vet ./...

## bench: 并发测试
.PHONY: bench
bench:
	go test ./...  -test.bench . -test.benchmem=true

## test: 单元测试全部测试代码
.PHONY: fmt
fmt:
	gofmt -w -l .

## lint: golangci-lint
.PHONY: lint
lint:
	golangci-lint cache clean
	golangci-lint run

## check: fmt lint vet
.PHONY: check
check: fmt lint vet

## build_linux: build_linux
.PHONY: build_linux
build_linux:
	rm -f main  && CGO_ENABLED=0 GOOS=linux go build -o  main -ldflags=${LDFLAGS}  main.go && go clean -cache

## build_mac_arm: build_mac_arm
.PHONY: build_mac_arm
build_mac_arm:
	rm -f main  && CGO_ENABLED=0 GOOS=darwin GOARCH=arm go build -o  main -ldflags=${LDFLAGS}  main.go && go clean -cache

## lint: install-lint
.PHONY: install-lint
install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2


## wire: service依赖自动组装生产
.PHONY: wire
wire:
	cd app && wire ./...


## http_serve: http服务启动
.PHONY: http_serve
http_serve:
	go run -ldflags=${LDFLAGS} main.go  http_serve


## help: Show this help info.
.PHONY: help
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'



%:
	@true