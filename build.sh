#!/bin/bash

docker run -it --rm -v $GOPATH/src/github.com/jpjohnson/go_example_cmdMenu -e GOOS=linux -e GOARCH=arm ubuntu `go build -v -o go_example_cmdMenu`
