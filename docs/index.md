## 通用web框架说明

从目录结构及使用模仿laravel，更符合phper的一款golang web框架

核心组件包: 
- gin (v1.9.1, web framework)
- go-redis  （v9.5.1, redis操作)
- gorm  (v1.25.9,mysql操作)
- cobra (v1.8.0, 命令行管理)
- zap (v1.27.0,日志输出)
- wire （v0.6.0, 依赖管理service)
- viper (v1.18.2, 配置管理)
- govalidator (v1.9.10, 数据验证包)
- 错误包 (https://github.com/marmotedu/errors)
- json包 （github.com/json-iterator/go v1.1.12) , 序列化是任何语言吃cpu的东西，使用更快的包。

使用文档
- [框架目录说明](框架目录说明.md)
- [结构分层](结构分层.md)
- [错误处理](错误处理.md)
- [请求验证器使用](请求验证器使用.md)
- [http错误码设计](http错误码设计.md)
- [依赖注入wire使用](依赖注入wire使用.md)
- [json包](json包.md)








