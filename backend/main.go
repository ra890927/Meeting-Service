package main

import (
	docs "meeting-center/docs"

	"meeting-center/src/clients"
	"meeting-center/src/middlewares"
	"meeting-center/src/presentations"
	"meeting-center/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.InitConfig()
	clients.InitDB()
}

// @title Meeting Center API
// @version 1.0
// @description This is a simple Meeting Center API

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RegisterMetricsMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		meeingPresentation := presentations.NewMeetingPresentation()
		roomPresentation := presentations.NewRoomPresentation()
		userPresentation := presentations.NewUserPresentation()
		filePresentation := presentations.NewFilePresentation()

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
		eg = v1.Group("/meeting")
		{
			eg.GET("/getAllMeetings", meeingPresentation.GetAllMeetings)
			eg.GET("/:id", meeingPresentation.GetMeeting)
			eg.GET("/getMeetingsByRoomIdAndDatePeriod", meeingPresentation.GetMeetingsByRoomIdAndDatePeriod)
			eg.GET("/getMeetingsByParticipantId", meeingPresentation.GetMeetingsByParticipantId)
			eg.POST("", middlewares.AuthRequire(), meeingPresentation.CreateMeeting)
			eg.PUT("", middlewares.AuthRequire(), meeingPresentation.UpdateMeeting)
			eg.DELETE("/:id", middlewares.AuthRequire(), meeingPresentation.DeleteMeeting)
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
		eg = v1.Group("/file")
		eg.Use(middlewares.AuthRequire())
		{
			eg.POST("", filePresentation.UploadFile)
			eg.GET("/:id", filePresentation.GetFile)
			eg.GET("/getFileURLsByMeetingID", filePresentation.GetFileURLsByMeetingID)
			eg.DELETE("/:id", filePresentation.DeleteFile)
		}
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
