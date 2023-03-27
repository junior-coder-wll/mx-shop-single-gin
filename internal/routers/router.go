//@Author: wulinlin
//@Description:
//@File:  router
//@Version: 1.0.0
//@Date: 2023/03/10 03:35

package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mx-shop-single-gin/internal/handler/naHandler"
	"mx-shop-single-gin/internal/middlewares"
)

func InitRouter(lg *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LoggerMiddleware(lg), middlewares.ErrorMiddleware())

	naRouter := r.Group("/api/v1/na")
	{
		naRouter.GET("/hello", naHandler.HelloController)
	}

	//AdminRouter :=

	return r
}
