#!/bin/bash
if [[ ! -z "$DOCKER_USER" || ! -z "$DOCKER_PWD" || ! -z "$DOCKER_REGISTRY"  ]]; then
  docker login -u="$DOCKER_USER" -p="$DOCKER_PWD" $DOCKER_REGISTRY
fi
echo "=> Download go deps..."
go get
echo "=> Start ! "
go run *.go
