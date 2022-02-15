#!/bin/bash

cd job-01/application || exit 
echo "Installing dependencies..."
go get -d -v ./.. && go install -v ./..
echo "Building application..."
go build -o calculator-app