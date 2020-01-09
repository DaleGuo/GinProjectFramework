# Go Web项目实践技术框架

**采用gin+gorm+mysql+casbin+nats(optional)+websocket(optional)**

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
    |--vendor
    |       |--vendor.json//项目依赖包配置
    |--web
    |	|--html
    |	|--static
    |	|	|--js
    |	|	|--css
    |	|	|--lib
    |--main.go //主程序
    |--ginFrameWork.sql //测试数据库脚本

# 项目配置
```
[webRoute]
port = :8081

[mysql]
user = root
password = adm123
address = 127.0.0.1:3306
dbName = ginFramework

[webSocket]
url = /processEdit
port = :8082

[Nats]
url = nats://localhost:4222
topic = natsTopic
```
    
# 启动web服务
```
gin.SetMode(gin.DebugMode)
router := gin.Default()
router.Static("../static", "./web/static")                             //静态资源
router.LoadHTMLFiles("./web/html/login.html", "./web/html/index.html") //静态页面
webRouter := router.Group("/ginFrameWork")
{
	webRouter.GET("/login", controller.LoginHtml)
	webRouter.GET("/index", controller.IndexHtml)
	webRouter.POST("/signIn", controller.SignIn)
	webRouter.POST("/signOut", controller.SignOut)
	webRouter.GET("/resource1", controller.GetResource1)
	webRouter.POST("/resource1", controller.PostResource1)
	webRouter.GET("/resource2", controller.GetResource2)
	webRouter.POST("/resource2", controller.PostResource2)
	//todo：其它路由
}
router.Run(port)
```

# session
```
store := cookie.NewStore([]byte("secret"))
router.Use(sessions.Sessions("mysession", store))
```

# 登录认证
```
router.Use(middleware.AuthenMiddleWare())

//认证中间件
func AuthenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session := sessions.Default(c); session.Get("hasSignIn") == "true" {
			c.Next()
			return
		}

		//登录页面和登录、退出操作跳过验证
		if url := c.Request.URL.String(); strings.HasPrefix(url, "/ginFrameWork/login") || strings.HasPrefix(url, "/ginFrameWork/signIn") || strings.HasPrefix(url, "/ginFrameWork/signOut") {
			c.Next()
			return
		}

		//静态资源文件跳过验证
		url := c.Request.URL.RequestURI()
		if strings.HasPrefix(url,"/static"){
			c.Next()
			return
		}

		c.HTML(http.StatusOK,"login.html",gin.H{"status":1,"message": "用户未登录"})
		c.Abort()
		return
	}
}
```


# 权限鉴定

**Casbin Access Control Model**

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

**Casbin Policy**

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

**权限鉴定中间件**

```
import _ "github.com/go-sql-driver/mysql" //一定要import

policyUrl,err:=global.GetAccessPolicyUrl()
a := gormadapter.NewAdapter("mysql",policyUrl, true)
e := casbin.NewEnforcer("./config/authz_model.conf", a)
e.LoadPolicy()//从DB加载策略
router.Use(middleware.AuthzMiddleWare(e))

//鉴权中间件
func AuthzMiddleWare(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//登录页面和登录、退出操作跳过验证
		if url := c.Request.URL.String(); strings.HasPrefix(url,"/ginFrameWork/login") || strings.HasPrefix(url,"/ginFrameWork/signIn") || strings.HasPrefix(url,"/ginFrameWork/signOut"){
			c.Next()
			return
		}

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := sessions.Default(c).Get("role")

		//静态资源文件跳过验证
		if strings.HasPrefix(obj,"/static"){
			c.Next()
			return
		}

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"status":1,"message": "用户无权限"})
			c.Abort()
		}
	}
}
```

**修改权限**

[API官方文档](https://github.com/casbin/gorm-adapter)
```
Enforcer.AddPolicy("student","/ginFrameWork/resource1","POST")//sub,obj,act
Enforcer.SavePolicy()

Enforcer.RemovePolicy("student","/ginFrameWork/resource1","POST")//sub,obj,act
Enforcer.SavePolicy()
```


# 部署并测试
1.  安装go 1.8环境
2.  安装包管理工具govendor `go get -u -v github.com/kardianos/govendor`
3.  进入ginFrameWork目录,输入`govendor sync`下载依赖包
3.  修改config/serverConfig.ini配置信息
4.  mysql执行ginFrameWork.sql脚本文件
5.  运行访问http://localhost:8081/ginFrameWork/login
6.  测试用户:

| userName | password | role |
| ------ | ------ | ------ |
| admin | admin | admin |
| professor | professor | professor |
| student | student | student |

# 第三方开源库
1. [Gin](https://github.com/gin-gonic/gin)：HTTP web framework.
2. [gin-contrib-sessions](https://github.com/gin-contrib/sessions):Gin middleware for session management
3. [gorm](https://github.com/jinzhu/gorm)：The fantastic ORM(Object Relational Mapping）library for Golang, aims to be developer friendlylibrary.
4. [casbin](https://github.com/jinzhu/gorm)：An authorization library that supports access control models like ACL, RBAC, ABAC in Golang.
5. [casbin-gorm-adapter](https://github.com/casbin/gorm-adapter)：Gorm adapter for Casbin.
6. [nats](https://github.com/nats-io/nats.go)：Golang client for NATS, the cloud native messaging system.
7. [gorilla-websocket](https://github.com/gorilla/websocket):A fast, well-tested and widely used WebSocket implementation for Go.
8. [go-ini](https://github.com/go-ini/ini):Package ini provides INI file read and write functionality in Go.
9. [go.uuid](https://github.com/satori/go.uuid):UUID package for Go.