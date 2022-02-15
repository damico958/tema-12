#!/bin/bash

cd application || exit 
echo "Installing dependencies..."
go get -d -v ./.. && go install -v ./..
echo "Building application..."
CGO_ENABLED=0 go build -o calculator-app
pwd 