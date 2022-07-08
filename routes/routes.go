package routes

import (
	"BaseWebProject/controller"
	"BaseWebProject/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", controller.UserRegister)
		v1.POST("user/login", controller.UserLogin)
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//任务操作
		}
	}
	return r
}