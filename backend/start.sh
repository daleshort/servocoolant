#!/bin/bash

export CC=aarch64-linux-gnu-gcc
GOARCH=arm64 GOOS=linux CGO_ENABLED=1 /usr/local/go/bin/go run  main.go