#!/bin/bash

docker run -it --rm -v $GOPATH/src/github.com/jpjohnson/go_example_cmdMenu -e GOOS=linux -e GOARCH=amd64 ubuntu `env GOOS=linux GOARCH=amd64 go build -v -o go_example_cmdMenu`
