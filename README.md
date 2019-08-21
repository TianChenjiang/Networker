# Networker



### 本地环境搭建

##### 必要依赖准备:

- go 11及以后版本
- Docker-compose &  docker
- node

##### 数据资源

在项目根目录下执行 `docker-compose up -d`, 即可在本机搭建 mysql 数据库.

数据库相关配置如下

| 属性               | 内容         |
| ------------------ | ------------ |
| Version 数据库版本 | Mysql 5.7.21 |
| Password 密码      | mysql        |
| Port 本机端口      | 33006        |

##### Go服务端开启

```bash
cd citicup-admin/app   							# 进入后端目录
export GOPROXY=https://goproxy.io 	# 代理配置
go build -o app 									  # 项目构建
./app
```

正常开启后端显示如下

```
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (2 handlers)
[GIN-debug] GET    /api/users                --> citicup-admin/internal/web.GetUserList (2 handlers)
2019/08/21 23:18:34 set up the container
```

此后可以在 `localhost:8080` 进行请求发送

