package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func RespuestaComentarioRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/respuestas-comentario")
	{
		grupo.GET("/", controllers.GetRespuestasComentario)
		grupo.GET("/:id", controllers.GetRespuestaComentarioByID)
		grupo.POST("/", controllers.CreateRespuestaComentario)
		grupo.PUT("/:id", controllers.UpdateRespuestaComentario)
		grupo.DELETE("/:id", controllers.DeleteRespuestaComentario)
	}
}