package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-lim/controller"
)

func InitApiRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	apiController := controller.NewApiController()
	router := r.Group("/api")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/list", apiController.GetApis)
		router.GET("/all/category/:roleId", apiController.GetAllApiGroupByCategoryByRoleId)
		router.POST("/create", apiController.CreateApi)
		router.PATCH("/update/:apiId", apiController.UpdateApiById)
		router.DELETE("/delete/batch", apiController.BatchDeleteApiByIds)
	}

	return r
}
