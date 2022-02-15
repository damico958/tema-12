#!/bin/bash

cd application || exit 
echo "Installing dependencies..."
go get -d -v ./.. && go install -v ./..
echo "Building application..."
go build -o calculator-app