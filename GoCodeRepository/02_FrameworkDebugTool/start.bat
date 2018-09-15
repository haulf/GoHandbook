@echo off

:: Copyright (C) 2017~2018 haulf

:: modify begin
set TAREGET_JAR=services.jar
set CLASS_ROOT_PATH=D:\Android_Studio_Debug_Space\workspace\FrameworkDebug\FrameworkDebug81\build\classes\java\main
:: modify end

set TOOL_PATH=D:\DebugSpace\FrameworkDebugTools
set CLASS_PATH=%TOOL_PATH%\ClassPath.txt

echo delete temp file, begin...
del %TOOL_PATH%\*.jar
del %TOOL_PATH%\tmp\*.jar
del %TOOL_PATH%\log
echo delete temp file, end...

@echo on

go run %TOOL_PATH%\UpdateJarAndDex.go -r %CLASS_ROOT_PATH% -c %CLASS_PATH% -j %TOOL_PATH%\debug_jar\%TAREGET_JAR% -o %TAREGET_JAR%
pause

rem echo.

rem adb root
rem adb wait-for-device

rem adb remount
rem adb wait-for-device

rem adb root
rem adb wait-for-device

rem echo.
rem adb push %TAREGET_JAR% /system/framework/

rem echo start to reboot
rem pause
rem adb reboot reboot

rem echo.
rem pause