#!/usr/bin/env bash

export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

APP_NAME=$1
PRJ_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && cd .. && pwd)
GO_VERSION=$(go env GOVERSION)
GIT_REPO_NAME=$(basename $(git rev-parse --show-toplevel))
GIT_REPO_URL=$(git remote get-url --push origin)
GIT_TAG=$(git describe --tags)
GIT_HASH=$(git rev-parse --short HEAD)
GIT_COMMIT_NUMBER=$(git rev-list --all --count)
APP_VERSION=$(git describe --tags)
BUILD_OS=$(go env GOHOSTOS)
BUILD_ARCH=$(go env GOHOSTARCH)
TARGET_OS=$(go env GOOS)
TARGET_ARCH=$(go env GOARCH)
BUILD_TIME=$(date +"%Y-%m-%d-%H:%M:%S")
BUILD_USER=$(git config user.email)
GO_MOD_NAME=$(go list -m)

# golang build args
GO_LDFLAGS="-s -w -X ${GO_MOD_NAME}/version.AppName=${GIT_REPO_NAME} \
            -X ${GO_MOD_NAME}/version.AppVersion=${APP_VERSION}  \
            -X ${GO_MOD_NAME}/version.GoVersion=${GO_VERSION} \
            -X ${GO_MOD_NAME}/version.GitRepoUrl=${GIT_REPO_URL} \
            -X ${GO_MOD_NAME}/version.GitTag=${GIT_TAG} \
            -X ${GO_MOD_NAME}/version.GitHash=${GIT_HASH} \
            -X ${GO_MOD_NAME}/version.BuildOS=${BUILD_OS} \
            -X ${GO_MOD_NAME}/version.BuildArch=${BUILD_ARCH} \
            -X ${GO_MOD_NAME}/version.TargetOS=${TARGET_OS} \
            -X ${GO_MOD_NAME}/version.TargetArch=${TARGET_ARCH} \
            -X ${GO_MOD_NAME}/version.BuildUser=${BUILD_USER} \
            -X ${GO_MOD_NAME}/version.BuildTime=${BUILD_TIME}"



function main() {
    echo "Building ${APP_NAME}"
    cd ${PRJ_DIR}
    [ -d "./bin" ] || mkdir -p ./bin
    # go build -ldflags "${GO_LDFLAGS}" -o ./bin ./...
    go build -ldflags "${GO_LDFLAGS}" -o ./bin/${APP_NAME}-${TARGET_OS}-${TARGET_ARCH}-${APP_VERSION} .
}

main
