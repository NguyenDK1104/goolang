package controller

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gindump "github.com/tpkeeper/gin-dump"
	"go-jwt/docs"
	"go-jwt/gorm/initializers"
	"go-jwt/gorm/service"
	"go-jwt/video/middlewares"
)

func init() {
	initializers.ConnectDb()
}

func Town() {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	docs.SwaggerInfo.Title = "Chanel-Smart"
	docs.SwaggerInfo.Description = "all-One-In"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9999"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"ptth"}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/get-all", service.GetAll())
		v1.GET("/get-by-id/:id", service.GetById())
		v1.POST("/create", service.Create())
		v1.GET("/get-by-name", service.GetByNameAndId())
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
