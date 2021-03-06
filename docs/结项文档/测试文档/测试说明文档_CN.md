# 测试说明文档

## 1. 测试简介

针对后端服务器的代码进行黑盒和白盒测试

## 2. 测试目的

- 通过对测试结果的分析，得到对软件质量的评价
- 评估测试执行和测试计划是否符合
- 分析系统存在的缺陷，为修复和预防bug提供建议

## 3. 测试工具

- Postman来进行服务的功能性测试
- Jmeter来进行接口压力测试

## 4. 测试流程

### 4.1 功能测试

#### 用户管理

| 功能                 | 基本要求                                          | 测试情况 | 通过(Y/N) |
| -------------------- | ------------------------------------------------- | -------- | --------- |
| 用户登录             | 认证成功后返回正确的token. 并且能够提供之后的授权 | 已实现   | Y         |
| 用户信息获取         | 用户获取个人信息, 还需要通过压力测试              | 已实现   | Y         |
| 修改用户信息         | 能够正常修改用户信息, 并且通过表单验证            | 已实现   | Y         |
| 用户标记关注公司     | 用户的公司关注列表正常添加                        | 已实现   | Y         |
| 用户获取关注公司列表 | 能够正常返回用户关注的公司列表，并且分页正常      | 已实现   | Y         |

#### 公司管理

| 功能         | 基本要求                                                   | 测试情况 | 通过(Y/N) |
| ------------ | ---------------------------------------------------------- | -------- | --------- |
| 获取公司信息 | 获取Tushare上的信息,并且正常显示                           | 已实现   | Y         |
| 公司模糊搜索 | 能够根据用户输入的关键字进行公司名称模糊搜索，并且分页正常 | 已实现   | Y         |

#### 股票信息管理

| 功能         | 基本要求                         | 测试情况 | 通过(Y/N) |
| ------------ | -------------------------------- | -------- | --------- |
| 获取股票信息 | 获取Tushare上的信息,并且正常显示 | 已实现   | Y         |

#### 公司财务信息

| 功能             | 基本要求                                                 | 测试情况 | 通过(Y/N) |
| ---------------- | -------------------------------------------------------- | -------- | --------- |
| 获取公司财务信息 | 根据公司的股票代码获取, 获取Tushare上的信息,并且正常显示 | 已实现   | Y         |

#### 股价平仓线预警概率

| 功能         | 基本要求                                                  | 测试情况 | 通过(Y/N) |
| ------------ | --------------------------------------------------------- | -------- | --------- |
| 获取爆仓概率 | 根据公司的股票代码和平仓线数据，应当返回爆仓概率 (30天内) | 已实现   | Y         |

### 4.2 性能测试

#### 测试参数说明

| 参数名         | 值              |
| -------------- | --------------- |
| 脚本循环次数   | 6000            |
| 并发用户数     | 5/10/50/100/200 |
| 真实客户端数量 | 1               |
| 脚本录制方法   | 自动            |
| 模拟路线类型   | 10/100M以太网   |

#### 性能测试具体结果

##### 执行效率

平均响应时间(s)：

| 并发用户数 | 登录  | 用户信息获取 | 获取股票信息 | 获取公司信息 |
| ---------- | ----- | ------------ | ------------ | ------------ |
| 5          | 0.195 | 0.258        | 0.421        | 0.033        |
| 10         | 0.294 | 0.213        | 0.293        | 0.057        |
| 50         | 2.502 | 1.400        | 1.501        | 2.714        |
| 100        | 5.228 | 3.775        | 6.841        | 1.812        |

### 性能分析

- 后端性能测试分析

根据JMeter聚合报告的结果可以看出，各后端api表现正常，在跑100个线程（用户）时，平均时间为9.862s，错误率在低并发(5-100)的情况下为0%

- 网页性能测试分析

根据JMeter聚合报告的结果可以看出，在跑200个线程（虚拟用户）时，当发出https请求数量为460左右，平均响应时间是1.452s而错误率仅是0.07%，符合预期的性能结果

![](https://i.loli.net/2019/09/14/sx8vbwlid3O1LSE.png)