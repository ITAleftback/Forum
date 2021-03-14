package routes

import (
	"forum_Anpw/controller"
	"forum_Anpw/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(c *gin.Engine) *gin.Engine {

	user:=c.Group("/user")
	{
		user.POST("/register",controller.Register)
		user.POST("/login",controller.Login)
		user.POST("/resetcode",controller.ResetCode)
		user.GET("/currentUser",middleware.AuthMiddleware(),controller.CurrentUser)
	}
	essay:=c.Group("/essay",middleware.AuthMiddleware())
	{
		essay.POST("/writeEssay",controller.WriteEssay)
		essay.POST("/currentEssay",controller.CurrentEssay)
		essay.POST("/yourEssay",controller.YourEssay)
		essay.POST("/delEssay",controller.DelEssay)
		essay.POST("/editEssay",controller.EditEssay)
	}
	comment:=c.Group("/comment",middleware.AuthMiddleware())
	{
		comment.POST("/essaycomment",controller.EssayComment)
		comment.POST("/currentComment",controller.CurrentComment)
		comment.POST("/delComment",controller.DelComment)
	}


	return c
}