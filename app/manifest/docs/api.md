# JD-WEB
[![Build Status](https://img.shields.io/badge/build-1.01-brightgreen)](https://travis-ci.org/pibigstar/go-todo)

> 此项目是模拟京东电商网站的一个后端平台，实现电商网站的基本功能。
>
> 因为前端队友跑路了，所以自己写的模板文件，写得太烂了，页面美观度保证不了,所以别介意页面了。。。

## 1.项目结构

<details>
<summary>展开查看</summary>
<pre><code>
    ├── app                        (项目文件)
        ├── api                    (api层)
        ├── global                 (全局对象)
        ├── internal               (内部逻辑)
            ├── consts             (常量定义)
            ├── middleware         (中间件层)
            ├── model              (模型层)
                ├── config         (配置文件模型)
                ├── request        (http请求模型)
                ├── response       (http响应模型)
            ├── service            (服务层)
        ├── manifest               (交付清单)
            ├── config             (配置管理)
            ├── docs               (swag接口文档)
            ├── docker             (镜像文件)
            ├── sql                (mysql初始化配置文件)
        ├── resource               (资源文件)
            ├── public             (静态资源)
            ├── template           (模板文件)
        ├── router                 (路由层)
    ├── boot                       (项目启动文件)
    ├── cert                       (证书目录)
    ├── log                        (日志目录)
    ├── utils                      (工具包)                    
</pre></code>
</details>

## 2. 使用技术

- [x] Swagger接口文档 实现 openapi v3规范
- [x] 基于Viper的配置文件管理系统
- [x] 基于Zap+Lumberjack的日志管理系统
- [x] 基于Nginx实现https全站加密, 静态文件代理
- [x] 使用mysql数据库和redis作服务器缓存
- [x] 基于singleflight包 进行并发控制防止缓存击穿
- [x] JWT+Oauth2.0的用户鉴权   (尚未实现两者整合)
- [x] 对用户和服务器之间传输的cookie,password进行SHA256,MD5加密
- [x] 实现支付宝支付接口
- [ ] 基于 ElasticSearch 实现网站搜索引擎

### 2.1 后端框架
使用`Gin`框架, golang 版本: `1.7.1`

##### 特性

Gin 是一个用 Go 编写的 HTTP web 框架。 它是一个类似于 [martini](https://github.com/go-martini/martini) 但拥有更好性能的 API 框架, 优于 [httprouter](https://github.com/julienschmidt/httprouter)，速度提高了近 40 倍。如果你需要极好的性能，使用 Gin 吧。

### 2.2 接口文档

使用`Swagger`自动化生成接口文档  [点击访问](`https://www.stellaris.wang/swagger/index.html`)

##### 特点

自动化生成RESTful API 文档

##### Swag设置

`./main.go`

`/app/api`

##### 生成文档

`/app/manifest/docs`

### 2.3 配置文件读取
使用`viper`框架 

##### 特点

- 找到和反序列化 JSON TOML YAML HCL JAVA配置文件
- 提供一个配置文件默认值和可选值的机制
- 提供重写配置值和Flag的可选值
- 提供系统的参数别名,解决对已有代码的侵入
- 轻松的辨别出用户输入值还是配置文件值

##### 配置模型

`/app/internal/model/config`

##### 配置文件

`/app/manifest/config/config.yaml`

### 2.3 日志输出
使用`zap`框架日志输出+`Lumberjack`进行日志分割

##### 特点

- 按不同的日志级别存储到不同的日志文件中(实现了`info` `debug` `warn` `error` `panic` 等各种日志级别)
- 控制最多保留的日志文件及保留最近多少天的日志文件(控制日志的最大磁盘使用量)
- 控制每个日志文件的大小

##### 详细配置

`/boot/zap.go`


### 2.4 Nginx
##### 特点

- 使用Nginx反向代理 提高访问速度，提供防火墙作用
- 使用Nginx托管网站证书，实现https全站加密
- 使用Nginx进行静态文件代理 提高访问速度

### 2.5 数据库

#### MySQL

##### 存储

- 页面信息

- 用户信息

#### Redis

##### 存储

- JWT
- Session
- 页面缓存

#### 并发控制

使用`singleflight`包进行并发控制防止缓存击穿

### 2.6 用户鉴权

#### JWT

- 提供令牌生效时间，过期时间，缓冲时间等配置选项
- 提供JWT黑名单功能

#### Oauth

实现GitHub第三方登录 (登录返回用户名验证是否有效,尚未实现于JWT整合)

### 2.7 安全性

对Cookie,Session及用户名密码进行加密，增加安全性

### 2.8 支付接口

实现了支付宝的支付功能

##### 沙盒账号

###### 账号: yntgxm3661@sandbox.com

###### 登录密码: 111111

###### 支付密码: 111111

## 3. 部署

### 3.1 编译成镜像
```bash
docker build -t web:v1 .
```
### 3.2 启动容器
```bash
docker run -d -p 8080:8080 -it --rm --name myweb web:v1
```

### 3.3 进入容器
```bash
docker exec -it myweb /bin/bash
```

### 3.4 删除镜像
```bash
# 停止容器
docker stop myweb
# 删除容器
docker rm myweb
# 删除镜像
docker rmi -f web:v1
```

