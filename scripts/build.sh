#!/bin/sh

set -e

BUILD_DATE=$(date -u --rfc-3339=seconds)
VERSION_PKG="github.com/yurykabanov/service-scaffold/pkg/version"

BUILD_FLAG="$VERSION_PKG.Build=${DRONE_COMMIT=unknown}"
VERSION_FLAG="$VERSION_PKG.Version=${DRONE_TAG=no_tag}"
BUILD_DATE_FLAG="$VERSION_PKG.BuildDate=$BUILD_DATE"

go mod download
go get -u github.com/swaggo/swag/cmd/swag

swag init --generalInfo ./cmd/server/main.go --output ./docs

go build \
  -ldflags="-X '$BUILD_FLAG' -X '$VERSION_FLAG' -X '$BUILD_DATE_FLAG'" \
  -o build/server \
  cmd/server/main.go
