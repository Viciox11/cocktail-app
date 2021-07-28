#!/bin/bash

echo "Insert build version (example 1.0):"
read version

echo "Preparing build: $version";
docker build -t ioartigiano-be:$version .
docker login -u ioartigiano
docker tag ioartigiano-be:$version ioartigiano/ioartigiano-be:$version
docker push ioartigiano/ioartigiano-be:$version

exit 0

