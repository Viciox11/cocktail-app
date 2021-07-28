#!/bin/sh

cd /opt/ioartigiano-be/cmd/server

pwd

go fmt && go build -o app *.go

ls

./app
