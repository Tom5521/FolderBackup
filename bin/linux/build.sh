#!/bin/bash
cd src
go mod tidy
go build -o ../vscodeback .