package router

import (
	"meeting-center/docs"
	"meeting-center/src/middlewares"
	"meeting-center/src/presentations"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Meeting Center API
// @version 1.0
// @description This is a simple Meeting Center API

func InitRouter() *gin.Engine {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	// CORS middleware
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RegisterMetricsMiddleware())

	v1 := r.Group("/api/v1")
	{
		userPresentation := presentations.NewUserPresentation()
		authPresentation := presentations.NewAuthPresentation()
		codePresentation := presentations.NewCodePresentation()
		meeingPresentation := presentations.NewMeetingPresentation()
		roomPresentation := presentations.NewRoomPresentation()
		filePresentation := presentations.NewFilePresentation()

		eg := v1.Group("/user")
		{
			eg.POST("", userPresentation.RegisterUser)
			eg.GET("/getAllUsers", userPresentation.GetAllUsers)
		}
		eg = v1.Group("/auth")
		{
			eg.POST("/login", authPresentation.Login)
			eg.POST("/logout", authPresentation.Logout)
			eg.GET("/whoami", middlewares.AuthRequire(), authPresentation.WhoAmI)
		}
		eg = v1.Group("/code")
		{
			// type routes
			eg.GET("/type/getAllCodeTypes", codePresentation.GetAllCodeTypes)
			eg.GET("/type/getCodeTypeByID", codePresentation.GetCodeTypeByID)

			// value routes
			eg.GET("/value/getCodeValueByID", codePresentation.GetCodeValueByID)

			sg := eg.Group("")
			sg.Use(middlewares.AuthRequire(), middlewares.AdminRequire())
			{
				// type routes
				sg.POST("/type", codePresentation.CreateCodeType)
				sg.PUT("/type", codePresentation.UpdateCodeType)
				sg.DELETE("/type", codePresentation.DeleteCodeType)

				// value routes
				sg.POST("/value", codePresentation.CreateCodeValue)
				sg.PUT("/value", codePresentation.UpdateCodeValue)
				sg.DELETE("/value", codePresentation.DeleteCodeValue)
			}
		}
		eg = v1.Group("/meeting")
		{
			eg.GET("/getAllMeetings", meeingPresentation.GetAllMeetings)
			eg.GET("/:id", meeingPresentation.GetMeeting)
			eg.GET("/getMeetingsByRoomIdAndDatePeriod", meeingPresentation.GetMeetingsByRoomIdAndDatePeriod)
			eg.GET("/getMeetingsByParticipantId", meeingPresentation.GetMeetingsByParticipantId)

			sg := eg.Group("")
			sg.Use(middlewares.AuthRequire())
			{
				sg.POST("", meeingPresentation.CreateMeeting)
				sg.PUT("", meeingPresentation.UpdateMeeting)
				sg.DELETE("/:id", meeingPresentation.DeleteMeeting)
			}
		}
		eg = v1.Group("/room")
		{
			eg.GET("/getAllRooms", roomPresentation.GetAllRooms)
			eg.GET("/:id", roomPresentation.GetRoomByID)
		}
		eg = v1.Group("/admin")
		eg.Use(middlewares.AuthRequire(), middlewares.AdminRequire())
		{
			sg := eg.Group("/user")
			{
				sg.PUT("", userPresentation.UpdateUser)
			}
			sg = eg.Group("/room")
			{
				sg.POST("", roomPresentation.CreateRoom)
				sg.PUT("", roomPresentation.UpdateRoom)
				sg.DELETE("/:id", roomPresentation.DeleteRoom)
			}
		}
		eg = v1.Group("/file")
		eg.Use(middlewares.AuthRequire())
		{
			eg.POST("", filePresentation.UploadFile)
			eg.GET("/:id", filePresentation.GetFile)
			eg.GET("/getFileURLsByMeetingID/:meeting_id", filePresentation.GetFileURLsByMeetingID)
			eg.DELETE("/:id", filePresentation.DeleteFile)
		}
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
