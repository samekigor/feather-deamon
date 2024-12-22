#!/bin/bash

source ./scripts/set-default-envs.sh

make build -o cmd/main.go

./bin/main
