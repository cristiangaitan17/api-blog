package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func NoticiaRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/noticias")
	{
		grupo.GET("/", controllers.GetNoticias)
		grupo.GET("/:id", controllers.GetNoticiaByID)
		grupo.POST("/", controllers.CreateNoticia)
		grupo.PUT("/:id", controllers.UpdateNoticia)
		grupo.DELETE("/:id", controllers.DeleteNoticia)
	}
}
