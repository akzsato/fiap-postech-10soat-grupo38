package route

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewDocRouter(router *gin.RouterGroup) {
	
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}