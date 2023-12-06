package main

import (
	"github.com/Ancordss/gpt-playground/configs"
	"github.com/Ancordss/gpt-playground/controller"
	"github.com/Ancordss/gpt-playground/middlewares"
	"github.com/Ancordss/gpt-playground/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
	Textservice     service.TextService        = service.New1()
	TextController  controller.TextController  = controller.New1(Textservice)
)

// func setupLogOutput() {
// 	f, err := os.Create("gin.log")
// 	if err != nil {
// 		fmt.Println("Error creating log file: ", err)
// 	}
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// }

func main() {
	// setupLogOutput()
	configs.SetupConfig()
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), middlewares.BasicAuth())

	if gin.Mode() == gin.TestMode {
		server.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(418, gin.H{
				"Content-Type": "application/json",
				"Accept":       "application/json",
				"status":       "ok",
				"mode":         "I don't exist in production",
			})
		})
	} else {
		server.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(418, gin.H{
				"Content-Type": "application/json",
				"Accept":       "application/json",
				"status":       "ok",
				"mode":         "I exist in production",
			})
		})
	}

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.POST("/ask", func(ctx *gin.Context) {
		ctx.JSON(200, TextController.Ask(ctx))
	})

	server.GET("/ask", func(ctx *gin.Context) {
		ctx.JSON(200, TextController.Gpt())
	})
	server.Run(":8080")
}
