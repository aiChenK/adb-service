@echo off

REM 进入当前目录
cd ./

REM 检查adb环境变量
adb version >nul 2>&1
if not %errorlevel% equ 0 (
    echo adb命令不存在，请设置环境变量!
    pause
    exit /b
)

REM 进入nginx目录
cd nginx

REM 启动 Nginx
start nginx

REM 等待几秒钟，以确保 Nginx 启动完成
timeout /t 5 >nul

REM 启动 Go 可执行文件
start /d "../" /B adb-service.exe

REM 关闭窗口时停止 Nginx
echo 关闭窗口以停止 Nginx...
echo 访问地址：http://localhost:8088
start http://localhost:8088
pause

REM 关闭Go可执行文件
taskkill /IM adb-service.exe /F

REM 停止 Nginx
nginx -s stop