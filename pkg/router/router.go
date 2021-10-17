package router

import (
	"io"
	"log"
	"os"

	"github.com/senthilsweb/notifier/controller"
	"github.com/senthilsweb/notifier/pkg/middleware"
	"github.com/senthilsweb/notifier/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Setup function
func Setup() *gin.Engine {
	r := gin.New()
	f, _ := os.Create(utils.AppExecutionPath() + "/" + os.Args[0] + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Println("Bootstrapping gin middlewares")
	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.GinContextToContextMiddleware())
	log.Println("Setting up routes")
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/notify/slack", controller.NotifySlack)
	r.POST("/api/notify/mailgun", controller.NotifyMailgun)
	r.POST("/api/pdf/export", controller.Export2PDF)
	r.POST("/api/png/export", controller.Export2PNG)
	r.POST("/api/redis/enqueue", controller.Enqueue)
	r.GET("/api/redis/dequeue/:key", controller.Dequeue)
	r.POST("/api/redis/hook/swissknife", controller.Swissknife)

	log.Println("Finished router setup")
	return r
}
