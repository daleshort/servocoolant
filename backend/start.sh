#!/bin/bash

export CC=aarch64-linux-gnu-gcc
GOARCH=arm64 GOOS=linux CGO_ENABLED=1 go run  main.go