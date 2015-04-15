#!/bin/bash

go build -o build/gogs-wrapper
docker build -t gogs .
