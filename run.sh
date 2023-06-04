#!/bin/sh

ARGS=$@

CGO_ENABLED=0 go run ./cmd/www "$ARGS"