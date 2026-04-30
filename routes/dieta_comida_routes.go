package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func DietaComidaRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/dieta-comidas")
	{
		grupo.GET("/", controllers.GetDietaComidas)
		grupo.GET("/:id", controllers.GetDietaComidaByID)
		grupo.POST("/", controllers.CreateDietaComida)
		grupo.PUT("/:id", controllers.UpdateDietaComida)
		grupo.DELETE("/:id", controllers.DeleteDietaComida)
	}
}
