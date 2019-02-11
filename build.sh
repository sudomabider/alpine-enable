#!/bin/sh

env GOOS=linux GOARCH=amd64 go build --ldflags="-s -w" -o build/enable.linux
