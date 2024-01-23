package demo

import (
	"github.com/lflxp/djangolang/examples/allinone/demo/service"
	"github.com/lflxp/djangolang/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterDemo(router *gin.Engine) {
	var authMiddleware = middlewares.NewGinJwtMiddlewares(middlewares.AllUserAuthorizator)
	demoGroup := router.Group("/apis/v1")
	demoGroup.Use(authMiddleware.MiddlewareFunc()) // 应用JWT认证中间件
	{
		demoGroup.GET("/demo", service.GetAllDemo)
		demoGroup.GET("/:id/demo", service.GetDemoById)
		demoGroup.POST("/demo", service.CreateDemo)
		demoGroup.DELETE("/:id/demo", service.DeleteDemo)
		demoGroup.PUT("/:id/demo", service.PutDemo)
	}
}

func RegisterVpn(router *gin.Engine) {
	var authMiddleware = middlewares.NewGinJwtMiddlewares(middlewares.AllUserAuthorizator)
	demoGroup := router.Group("/apis/v1")
	demoGroup.Use(authMiddleware.MiddlewareFunc()) // 应用JWT认证中间件
	{
		demoGroup.GET("/vpn", service.GetAllVpn)
		demoGroup.GET("/:id/vpn", service.GetVpnById)
		demoGroup.POST("/vpn", service.CreateVpn)
		demoGroup.DELETE("/:id/vpn", service.DeleteVpn)
		demoGroup.PUT("/:id/vpn", service.PutVpn)
	}
}
