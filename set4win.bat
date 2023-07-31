@echo off
copy bin\linux\* .
echo Compiling...
call build.sh
echo Compiled!
pause