#!/bin/bash

# 启动第一个可执行程序
/app/adb_server &

# 启动第二个可执行程序
nginx -g "daemon off;"