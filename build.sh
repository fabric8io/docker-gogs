#!/bin/bash

go build -o build/gogs-wrapper
docker build --no-cache -t gogs .
