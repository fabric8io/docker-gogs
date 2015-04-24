#!/bin/bash

go get github.com/tools/godep
godep go build -o build/gogs-wrapper
docker build -t gogs .
