package main

import (
	docs "meeting-center/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"meeting-center/src/presentations"
	"meeting-center/src/utils"
)

func main() {
	// connect to the database
	err := utils.InitDB()
	if err != nil {
		panic(err)
	}
	// Create a new presentation

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/User")
		{
			userPresentation := presentations.NewUserPresentation()
			eg.POST("", userPresentation.RegisterUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
