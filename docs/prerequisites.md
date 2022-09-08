# Prerequisites

You should have the following tools installed.

**golang** installed via **brew**

## tools

### brew

```shell
# github CLI
brew install gh
# for mac, use brew to install protobuf
# buf: proto tool https://buf.build/docs/tour-1
brew install bufbuild/buf/buf
brew install protobuf
brew install grpcurl
# bloomrpc is a UI client for gRPC (optional)
# install `bloomrpc` via `brew` into ~/Applications)
brew cask install --appdir=~/Applications bloomrpc
# gRPC mock server for testing
yarn global add bloomrpc-mock
```

### golang

```shell
# Update outdated Go dependencies interactively
# Usage: go-mod-upgrade ./...
go install github.com/oligot/go-mod-upgrade@main
# for static check/linter
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# linter and tool for proto files
# *** (if you use brew to install buf, skip next line) ***
go install github.com/bufbuild/buf/cmd/buf@latest
# go lang  build/publish/deploy tool (optional)
go install github.com/google/ko/cmd/ko@latest

# fetch protoc plugins into $GOPATH
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/srikrsna/protoc-gen-gotag@latest

# Installing PGV can currently only be done from source: 
# from user's home directory, run
export GOPATH = $(go env GOPATH)
go get -d github.com/envoyproxy/protoc-gen-validate
cd $GOPATH/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7
git pull
make build
```

### VSCode 

Setting VSCode with [1.18](https://github.com/golang/vscode-go/blob/master/docs/advanced.md)
