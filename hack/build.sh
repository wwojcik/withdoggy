#!/bin/sh
export GIT_COMMIT=$(git rev-list -1 HEAD)
go build -a -installsuffix cgo -ldflags "-w -s -X main.Version=$GIT_COMMIT" -o bin/api-gateway cmd/api-gateway/main.go