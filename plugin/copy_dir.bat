
@REM source_dir="E:\code\gin-vue-admin\server"
@REM target_dir="E:\code\gin_web_frame"

@echo off
rem 定义源文件夹路径，你需要将其替换为实际的源文件夹路径
set "source_folder=E:\code\gin-vue-admin\server"
rem 定义目标文件夹路径，你需要将其替换为实际的目标文件夹路径
set "target_folder=E:\code\gin_web_frame"

rem 检查目标文件夹是否存在，如果不存在则创建
if not exist "%target_folder%" (
    mkdir "%target_folder%"
)

rem 遍历源文件夹下的所有子文件夹
for /d %%i in ("%source_folder%\*") do (
    rem 获取当前子文件夹的名称
    set "folder_name=%%~ni"
    rem 启用延迟环境变量扩展
    setlocal enabledelayedexpansion
    rem 在目标文件夹中创建对应的空文件夹
    mkdir "%target_folder%\!folder_name!"
    endlocal
)