#!/bin/sh

set -e

go mod download
go get -u github.com/swaggo/swag/cmd/swag

swag init --generalInfo ./cmd/server/main.go --output ./docs

go test ./...
