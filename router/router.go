package router

import (
	"github.com/gin-gonic/gin"
	"GinProjectFramework/controller"
	"GinProjectFramework/global"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/casbin/casbin"
	"GinProjectFramework/middleware"
	"github.com/casbin/gorm-adapter"
)

func Route() error {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//登录认证
	router.Use(middleware.AuthenMiddleWare())

	//权限认证
	policyUrl,err:=global.GetAccessPolicyUrl()
	a := gormadapter.NewAdapter("mysql",policyUrl, true)
	e := casbin.NewEnforcer("../config/authz_model.conf", a)
	e.LoadPolicy()//从DB加载策略
	router.Use(middleware.AuthzMiddleWare(e))

	router.Static("../static", "./web/static")                             //静态资源
	router.LoadHTMLFiles("./web/html/login.html", "./web/html/index.html") //静态页面

	webRouter := router.Group("/ginFrameWork")
	{
		webRouter.GET("/login", controller.LoginHtml)
		webRouter.GET("/index", controller.IndexHtml)
		webRouter.POST("/signIn", controller.SignIn)

		//todo：其它路由
	}

	//监听端口
	port, err := global.GetWebRoutePort()
	if err != nil {
		log.Println("http端口号获取失败", err)
		return err
	}

	router.Run(port)
	log.Println("http服务启动成功")
	return nil
}
