@echo off

REM ���뵱ǰĿ¼
cd ./

REM ���adb��������
adb version >nul 2>&1
if not %errorlevel% equ 0 (
    echo adb������ڣ������û�������!
    pause
    exit /b
)

REM ����nginxĿ¼
cd nginx

REM ���� Nginx
start nginx

REM �ȴ������ӣ���ȷ�� Nginx �������
timeout /t 5 >nul

REM ���� Go ��ִ���ļ�
start /d "../" /B adb-service.exe

REM �رմ���ʱֹͣ Nginx
echo �رմ�����ֹͣ Nginx...
echo ���ʵ�ַ��http://localhost:8088
start http://localhost:8088
pause

REM �ر�Go��ִ���ļ�
taskkill /IM adb-service.exe /F

REM ֹͣ Nginx
nginx -s stop