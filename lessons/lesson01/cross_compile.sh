#!/bin/bash

# See more: https://golang.org/doc/install/source#environment

#echo "build http-linux"
#GOOS=linux GOARCH=amd64 go build -o bin/http-linux http.go

echo "build http-osx"
GOOS=darwin GOARCH=amd64 go build  -o bin/http-osx http.go

#echo "build http-windows"
#GOOS=windows GOARCH=386 go build -o bin/http-win http.go
