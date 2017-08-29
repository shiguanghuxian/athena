@echo off
@rem #"Makefile" for Windows monitor projects.
@rem #Copyright (c) http://www.atisafe.com/, 2015. All rights reserved.
@rem
SETLOCAL

@rem ###############
@rem # PRINT USAGE #
@rem ###############

if [%1]==[] goto usage

@rem ################
@rem # SWITCH BLOCK #
@rem ################

@rem # make build
if /I [%1]==[build] call :build

@rem # make clean
if /I [%1]==[clean] call :clean

goto :eof

@rem #############
@rem # FUNCTIONS #
@rem #############
:build
for /f "delims=" %%i in ('go version') do (set go_version=%%i)
for /f "delims=" %%i in ('git rev-parse HEAD') do (set git_hash=%%i)
@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATE% %TIME%' -X 'main.GO_VERSION=%go_version%' -X main.GIT_HASH=%git_hash%" -o ./build/monitorevent.exe ./apps/monitorevent
@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATE% %TIME%' -X 'main.GO_VERSION=%go_version%' -X main.GIT_HASH=%git_hash%" -o ./build/monitorserver.exe ./apps/monitorserver
@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATE% %TIME%' -X 'main.GO_VERSION=%go_version%' -X main.GIT_HASH=%git_hash%" -o ./build/monitorweb.exe ./apps/monitorweb
@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATE% %TIME%' -X 'main.GO_VERSION=%go_version%' -X main.GIT_HASH=%git_hash%" -o ./build/monitorcron.exe ./apps/monitorcron
@go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATE% %TIME%' -X 'main.GO_VERSION=%go_version%' -X main.GIT_HASH=%git_hash%" -o ./build/monitorstorage.exe ./apps/monitorstorage
exit /B %ERRORLEVEL%

:clean
@del /S /F /Q "build\monitorevent.exe" "build\monitorserver.exe" "build\monitorweb.exe" "build\monitorcron.exe" "build\monitorstorage.exe"
@del /S /F /Q "build\logs\*.log"
exit /B 0

:usage
@echo Usage: %0 ^[ build ^| clean ^]
exit /B 1