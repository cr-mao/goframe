##  go web 脚手架框架

快速、简单、对phper友好

基于gin的脚手架框架，复制过来快速改改就能开发。 

go version:  1.22，当然你可以降低版本，这都没问题。

###  Useage

```shell
go mod download
go run main.go --help 

## 项目依赖二进制工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
go install github.com/google/wire/cmd/wire@v0.6.0


# 启动http 服务接受 消息
go run main.go http_serve --env=local


 # 另外Makefile 里面有一些命令
 make help
 # 如 代码检测
 make check 
```

#### docker 初始化demo表

```shell
docker build -t mysql-for-go:v1 -f mysqlDockerfile .
docker run -itd --name mysqlforgodemo -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql-for-go:v1
# mysql -h 127.0.0.1 -u root -P 3308 -p
```





### 相关文档

- [框架目录说明](docs/框架目录说明.md)
- [结构分层](docs/结构分层.md)
- [错误处理](docs/错误处理.md)
- [请求验证器使用](docs/请求验证器使用.md)
- [http错误码设计](docs/http错误码设计.md)
- [依赖注入wire使用](docs/依赖注入wire使用.md)
- [json包](docs/json包.md)











