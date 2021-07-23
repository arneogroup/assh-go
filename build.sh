#!/usr/bin/env bash

mkdir -p builds/{windows,linux}_amd64
GOOS=windows GOARCH=amd64 go build -o builds/windows_amd64/
GOOS=linux GOARCH=amd64 go build -o builds/linux_amd64/
