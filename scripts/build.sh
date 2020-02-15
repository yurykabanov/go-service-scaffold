#!/bin/sh

BUILD_DATE=$(date -u --rfc-3339=seconds)
VERSION_PKG="github.com/yurykabanov/service-scaffold/pkg/version"

BUILD_FLAG="$VERSION_PKG.Build=${DRONE_COMMIT=unknown}"
VERSION_FLAG="$VERSION_PKG.Version=${DRONE_TAG=no_tag}"
BUILD_DATE_FLAG="$VERSION_PKG.BuildDate=$BUILD_DATE"

# TODO: dependencies?
#  go download ?

# TODO: fix paths?
# this works, but take so long to complete:
#   swag init -g ./cmd/server/main.go -o ../../docs -d ../../ --parseDependency
go generate ./cmd/server/main.go

go build \
  -ldflags="-X '$BUILD_FLAG' -X '$VERSION_FLAG' -X '$BUILD_DATE_FLAG'" \
  -o build/server \
  cmd/server/main.go
