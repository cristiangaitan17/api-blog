package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func CategoriaRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/categorias")
	{
		grupo.GET("/", controllers.GetCategorias)
		grupo.GET("/:id", controllers.GetCategoriaByID)
		grupo.POST("/", controllers.CreateCategoria)
		grupo.PUT("/:id", controllers.UpdateCategoria)
		grupo.DELETE("/:id", controllers.DeleteCategoria)
	}
}

