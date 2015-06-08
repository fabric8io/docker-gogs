#!/bin/bash

set -euf -o pipefail

go get github.com/tools/godep
godep go build -tags="sqlite" -a -o build/gogs-wrapper
docker build -t gogs .
