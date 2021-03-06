## Citicup Admin 微服务部分

### 1.负责资源内容

公司信息资源



### 2.部署环境说明

| 属性         | 内容                  |
| ------------ | --------------------- |
| 镜像名称     | citicup/citicup-admin |
| 内网占用端口 | 8080                  |
| 外网网桥     | citicup               |



### 3.编码说明

#### 3.1 Restful

- 使用HTTP方法类型进行资源状态改变的标识，以资源本身的名词作为资源标识
- 异常统一处理

#### 3.2 Error 与 Panic处理说明

考虑到Go本身对于异常和错误的支持欠缺，现对可能发生的异常和处理进行如下分层处理

| Layer            | Error & Panic                                  |
| ---------------- | ---------------------------------------------- |
| Dao 数据访问     | 资源缺失 (404<br>资源已存在<br>数据连接件超时  |
| Service 业务逻辑 | 权限(401)                                      |
| Web 网络控制     | 请求数据有误：格式错误(400)<br>网关拦截器(401) |

