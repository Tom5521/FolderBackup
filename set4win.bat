@echo off
copy bin\windows\* .
echo Compiling...
call build.bat
echo Compiled!
pause
