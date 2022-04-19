package router

import (
	"http_req/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.GetUrl)
	router.GET("/:id", controller.GetUrlID)
	router.POST("/", controller.PostUrl)
	return router
}
