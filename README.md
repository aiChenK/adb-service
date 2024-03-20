# adb-service

## Project setup
```
go mod tidy
```

### Compiles for development
```
go run .
```

### Compiles and minifies for production
```
go build -o adb-service.exe
```

### Build Local-Only
```
goreleaser release --snapshot --clean
```


## 更新日志

### 2024-03-20

* goreleaser构建支持

### 2024-03-19

* 静态资源合并，不再依赖nginx代理
* adb环境检查支持
* 启动后自动打开浏览器网页

### 2024-03-18

* 支持mac获取ip地址

### 2024-03-14

* 增加获取以太网ip接口
