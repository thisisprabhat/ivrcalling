@echo off
REM Q&I IVR API Server Startup Script
echo Starting Q&I IVR API Server...
echo.

REM Check if .env file exists
if not exist .env (
    echo WARNING: .env file not found!
    echo Please copy .env.example to .env and configure your Twilio credentials.
    echo.
    pause
    exit /b 1
)

REM Start the server
echo Server starting on port 8080...
echo Press Ctrl+C to stop the server
echo.
.\bin\ivr-api.exe
