package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tholgbg/golang-gin-poc/controller"
	"gitlab.com/tholgbg/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(300, videoController.Save(ctx))
	})

	server.Run(":8080")
}
