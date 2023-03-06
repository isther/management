package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/management/middleware"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(STATUS_SUCCESS, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/user/register", NewUserApi().Register)
	router.POST("/api/user/login", NewUserApi().Login)

	var jwtAuth = router.Group("/api", middleware.JWTAuth())
	{
		jwtAuth.POST("/hello", func(ctx *gin.Context) {
			token := ctx.GetHeader("Authorization")
			claims, err := middleware.ParseMapClaimsJwt(token)
			if err != nil {
				ctx.JSON(STATUS_FAILED, gin.H{
					"msg": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"msg": "hello, " + claims["username"].(string)})
		})

		jwtAuth.POST("/config/add", NewAdminApi().Add)         //添加系统设置
		jwtAuth.GET("/config/get", NewAdminApi().Get)          //获取系统设置
		jwtAuth.DELETE("/config/delete", NewAdminApi().Delete) //删除系统设置

		jwtAuth.POST("/item/create", NewItemApi().Create)   // 添加物料信息
		jwtAuth.GET("/item/get", NewItemApi().Query)        // 获取物料信息
		jwtAuth.PUT("/item/update", NewItemApi().Update)    // 修改物料信息
		jwtAuth.DELETE("/item/delete", NewItemApi().Delete) // 删除物料信息

		jwtAuth.POST("/inbound/create", NewDetailsApi().CreateInbound) // 添加入库记录
		jwtAuth.GET("/inbound/get", NewDetailsApi().QueryAllInbound)   // 获取入库记录
		// jwtAuth.PUT("/inbound/update", api.NewDetailsApi().UpdateInbound)    // 更新入库记录
		jwtAuth.DELETE("/inbound/delete", NewDetailsApi().DeleteInbound) // 删除入库记录

		jwtAuth.POST("/outbound/create", NewDetailsApi().CreateOutbound) // 添加出库记录
		jwtAuth.GET("/outbound/get", NewDetailsApi().QueryAllOutbound)   // 获取出库记录
		// jwtAuth.PUT("/outbound/update", api.NewDetailsApi().UpdateOutbound)    // 更新出库记录
		jwtAuth.DELETE("/outbound/delete", NewDetailsApi().DeleteOutbound) // 删除出库记录

		jwtAuth.POST("/query_by_timestamp", NewDetailsApi().QueryByTimestamp)
	}
	return router
}
