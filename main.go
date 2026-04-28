package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/routes"
)

func main() {
	config.InitDB()

	router := gin.Default()

	// Registrar rutas
	routes.CategoriaRoutes(router)
	routes.NoticiaRoutes(router)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}