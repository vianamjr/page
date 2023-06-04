#!/bin/sh

VERSION=$(./scripts/version.sh)
FLAGS=--ldflags="-X 'github.com/vianamjr/page/internal/www.version=${VERSION}'"

echo "Build version ${VERSION}"

echo 'Build for the current os and platform'
CGO_ENABLED=0 go build -o bin/www.current "$FLAGS" ./cmd/www

echo 'Build for linux amd'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/www.linux.amd64-"$VERSION" "$FLAGS" ./cmd/www
