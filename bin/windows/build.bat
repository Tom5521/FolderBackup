@echo off
cd src
go mod tidy
go build -o ..\vscodeback.exe .
pause