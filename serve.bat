@ECHO OFF

CD /D %~dp0

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::����ϵͳ����

SET GIN_MODE=debug

IF EXIST D:\RunTime\go\runtime.bat (
    CALL D:\RunTime\go\runtime set "%~n0"
)

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /D %~dp0

go run main.go

IF "%1" == "" CMD /K
