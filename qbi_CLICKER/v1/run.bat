@echo off
cd %~dp0
echo Compiling qbiGO AUTO Clicker...
go build -o run.exe run.go

if %ERRORLEVEL% neq 0 (
    echo Compilation failed!
    exit /b %ERRORLEVEL%
)

echo Starting Flask apps via Go...
start run.exe

:: Introduce a delay to ensure the server starts
timeout /t 1 > nul

echo Opening the server in the browser...
start "" http://127.0.0.1:8090

exit
