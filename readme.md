## gin Api 项目快速搭建

### 部署文档
```text

git clone https://github.com/dvliwei/go.gin.base.git go.gin.base

```

### 配置文件所在目录
```text
cp .env.example  .env

```

### 架构说明
```text
conf 目录为配置文件 通过读取.env文件到配置信息进行实现
http/controller  目录为controller 
http/middleware  中间键
modules/model    数据模型
modules/Repositories 数据处理
routers  路由
cache  存放日志
assets  静态文件夹
```

 
### 启动项目 自动生成文档
```text

go run main.go

```
###(完毕)
