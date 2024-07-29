#!/usr/bin/env bash

ls -la ws
rm -rf ws
ls -la ws
GOOS=linux GOARCH=amd64 go build -o ws main.go
scp -r ./ws hugo-docker.bogdi.xyz:~/
