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
		roomPresentation := presentations.NewRoomPresentation()
		userPresentation := presentations.NewUserPresentation()

		eg := v1.Group("/user")
		{
			eg.POST("", userPresentation.RegisterUser)
			eg.GET("/getAllUsers", userPresentation.GetAllUsers)
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
			eg.GET("/type/getCodeTypeByID", codePresentation.GetCodeTypeByID)
			eg.POST("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.CreateCodeType)
			eg.PUT("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.UpdateCodeType)
			eg.DELETE("/type", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.DeleteCodeType)

			// value routes
			eg.GET("/value/getCodeValueByID", codePresentation.GetCodeValueByID)
			eg.POST("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.CreateCodeValue)
			eg.PUT("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.UpdateCodeValue)
			eg.DELETE("/value", middlewares.AuthRequire(), middlewares.AdminRequire(), codePresentation.DeleteCodeValue)
		}
		eg = v1.Group("/room")
		{
			eg.GET("/getAllRooms", roomPresentation.GetAllRooms)
			eg.GET("/:id", roomPresentation.GetRoomByID)
		}
		eg = v1.Group("/admin")
		eg.Use(middlewares.AuthRequire(), middlewares.AdminRequire())
		{
			sub := eg.Group("/user")
			{
				sub.PUT("", userPresentation.UpdateUser)
			}
			sub = eg.Group("/room")
			{
				sub.POST("", roomPresentation.CreateRoom)
				sub.PUT("", roomPresentation.UpdateRoom)
				sub.DELETE("/:id", roomPresentation.DeleteRoom)
			}
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
