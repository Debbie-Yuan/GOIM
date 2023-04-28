package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"im/docs"
	"im/service"
)

func Router() *gin.Engine {
	router := gin.Default()
	// 测试
	router.GET("/index", service.IndexHandler)
	router.GET("/user/all", service.GetAllUsers)

	// Swagger doc integration
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
