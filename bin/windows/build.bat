@echo off
cd src
go mod tidy
go build -o ..\folderbackup.exe .
pause