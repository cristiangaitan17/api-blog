package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func ArticuloSeccionRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/articulos-secciones")
	{
		grupo.GET("/", controllers.GetArticuloSecciones)
		grupo.GET("/:id", controllers.GetArticuloSeccionByID)
		grupo.POST("/", controllers.CreateArticuloSeccion)
		grupo.PUT("/:id", controllers.UpdateArticuloSeccion)
		grupo.DELETE("/:id", controllers.DeleteArticuloSeccion)
	}
}
