package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/controllers"
)

func NutricionRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/nutricion")
	{
		grupo.GET("/", controllers.GetNutricion)
		grupo.GET("/:id", controllers.GetNutricionByID)
		grupo.POST("/", controllers.CreateNutricion)
		grupo.PUT("/:id", controllers.UpdateNutricion)
		grupo.DELETE("/:id", controllers.DeleteNutricion)
	}
}