#!/bin/bash

export CC=aarch64-linux-gnu-gcc
GOARCH=arm64 GOOS=linux CGO_ENABLED=1 GOMODCACHE=/home/dale/go/pkg/mod /usr/local/go/bin/go build -o servocoolant  main.go