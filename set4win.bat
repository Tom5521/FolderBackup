@echo off
copy bin\windows\* .
echo Compiling...
call build.sh
echo Compiled!
pause
