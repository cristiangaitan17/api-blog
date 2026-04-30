package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func ComentarioComunidadRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/comentarios-comunidad")
	{
		grupo.GET("/", controllers.GetComentariosComunidad)
		grupo.GET("/:id", controllers.GetComentarioComunidadByID)
		grupo.POST("/", controllers.CreateComentarioComunidad)
		grupo.PUT("/:id", controllers.UpdateComentarioComunidad)
		grupo.DELETE("/:id", controllers.DeleteComentarioComunidad)
	}
}
