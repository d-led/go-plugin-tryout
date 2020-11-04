@echo off

echo --== compiling main ==--
del /s randomizer_rpc_plugin.exe randomizer_plugin.dll
go build -o randomizer_demo.exe .
if %errorlevel% neq 0 exit /b %errorlevel%

echo --== running the demo without the plugin built ==--
randomizer_demo.exe

echo --== compiling the rpc plugin ==--
go build -o randomizer_rpc_plugin.exe github.com/d-led/go-plugin-tryout/randomizer_rpc
if %errorlevel% neq 0 exit /b %errorlevel%

@REM -buildmode=plugin not supported on windows/amd64
@REM echo --== compiling the native plugin ==--
@REM go build -buildmode=plugin -o randomizer_plugin.dll github.com/d-led/go-plugin-tryout/randomizer_native
@REM if %errorlevel% neq 0 exit /b %errorlevel%

echo --== running the demo with the plugins built ==--
randomizer_demo.exe
if %errorlevel% neq 0 exit /b %errorlevel%
