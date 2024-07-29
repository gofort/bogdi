#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o ws main.go
scp -r ./ws hugo-docker.bogdi.xyz:~/
