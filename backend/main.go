package main

import (
	docs "meeting-center/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"meeting-center/src/middlewares"
	"meeting-center/src/presentations"
	"meeting-center/src/utils"
)

// @title Meeting Center API
// @version 1.0
// @description This is a simple Meeting Center API

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
		eg := v1.Group("/user")
		{
			userPresentation := presentations.NewUserPresentation()
			eg.POST("", userPresentation.RegisterUser)
			eg.PUT("", userPresentation.UpdateUser)
		}
		eg = v1.Group("/auth")
		{
			authPresentation := presentations.NewAuthPresentation()
			eg.POST("/login", authPresentation.Login)
			eg.POST("/logout", authPresentation.Logout)
			eg.GET("/whoami", middlewares.AuthRequire(), authPresentation.WhoAmI)
		}
		eg = v1.Group("/code")
		{
			codePresentation := presentations.NewCodePresentation()
			// type routes
			eg.GET("/type/getAllCodeTypes", codePresentation.GetAllCodeTypes)
			eg.POST("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.CreateCodeType)
			eg.PUT("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.UpdateCodeType)
			eg.DELETE("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.DeleteCodeType)

			// value routes
			eg.POST("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.CreateCodeValue)
			eg.PUT("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.UpdateCodeValue)
			eg.DELETE("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.DeleteCodeValue)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
