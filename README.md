# Go Web项目实践技术框架

**采用gin+gorm+mysql+casbin+nats(optional)+websocket(optional)**

# 第三方开源库
1. [Gin](https://github.com/gin-gonic/gin)：HTTP web framework.
2. [gin-contrib-sessions](https://github.com/gin-contrib/sessions):Gin middleware for session management
3. [gorm](https://github.com/jinzhu/gorm)：The fantastic ORM(Object Relational Mapping）library for Golang, aims to be developer friendlylibrary.
4. [casbin](https://github.com/jinzhu/gorm)：An authorization library that supports access control models like ACL, RBAC, ABAC in Golang.
5. [nats](https://github.com/nats-io/nats.go)：Golang client for NATS, the cloud native messaging system.
6. [gorilla-websocket](https://github.com/gorilla/websocket):A fast, well-tested and widely used WebSocket implementation for Go.
7. [go-ini](https://github.com/go-ini/ini):Package ini provides INI file read and write functionality in Go.
8. [go.uuid](https://github.com/satori/go.uuid):UUID package for Go.

# 项目目录
    |--config //配置相关
    |	|--authz_model.conf //casbin鉴权模型配置
    |	|--serverConfig.ini //ini配置
    |--controller //控制器层，验证提交的数据，将验证完成的数据传递给service层
    |	|--controller.go //页面发布
    |	|--userController.go //用户对象相关操作
    |--db //数据库操作层
    |	|--mysql.go //数据库初始化
    |	|--userModel.go //用户表相关操作
    |--global //公共
    |	|--config.go //获取配置信息
    |--log //日志
    |--middleware //中间件
    |	|--authe.go //登录认证
    |	|--authz.go //权限鉴定
    |--otherInterface //其它
    |	|--nats
    |	|	|--nats.go //初始化和操作封装
    |	|	|--natsModel.go //nats模型
    |	|--websocket
    |	|	|--ws.go //初始化和操作封装
    |	|	|--wsModel.go //ws模型
    |--router //路由
    |	|--router.go //设置路由
    |--service //业务层，只完成业务逻辑的开发，不进行操作数据库
    |	|--userService.go //用户对象相关服务
    |--util //工具类
    |	|--unique.go //唯一标识生成
    |--web
    |	|--html
    |	|--static
    |	|	|--js
    |	|	|--css
    |	|	|--lib
    |--main.go //主程序
    |--ginFrameWork.sql //测试数据库脚本
    
# casbin权限鉴定配置

1.  Access Control Model

定义在/config/authz_model.conf中，配置匹配规则
```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
```

2.  Policy

定义在mysql的casbin_rule表中，设置用户权限

| p_type | v0 | v1 | v2 | v3 | v4 | v5 |
| ------ | ------ | ------ | ------ | ------ | ------ | ------ |
| p | admin | /ginFrameWork/* | (GET)(POST) |  |  |
| p | admin | /ginFrameWork/index | GET |  |  |
| p | professor | /ginFrameWork/resource1 | GET |  |  |
| p | professor | /ginFrameWork/resource2 | (GET)(POST) |  |  |
| p | professor | /ginFrameWork/index | GET |  |  |
| p | student | /ginFrameWork/resource1 | GET |  |  |
| p | student | /ginFrameWork/resource2 | GET |  |  |
| p | student | /ginFrameWork/index | GET |  |  |

第三行数据表示professor角色用户对/ginFrameWork/resource1资源只拥有GET权限

关于model和policy请参照[casbin](https://github.com/casbin/casbin)官方文档

# 部署并测试
1.  安装go环境
2.  下载工程，对缺少的第三方库使用`go get`获取，例如`go get github.com/gin-gonic/gin`
3.  修改config/serverConfig.ini配置信息
4.  mysql执行ginFrameWork.sql脚本文件
5.  运行访问http://localhost:8081/ginFrameWork/login
6.  测试用户:

| userName | password | role |
| ------ | ------ | ------ |
| admin | admin | admin |
| professor | professor | professor |
| student | student | student |