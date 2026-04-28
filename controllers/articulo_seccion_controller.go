package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetArticuloSecciones obtiene todas las secciones de artículos
func GetArticuloSecciones(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, COALESCE(articulo_id, 0), COALESCE(titulo_seccion, ''), 
		       COALESCE(contenido, ''), COALESCE(imagen_url, ''), COALESCE(orden, 0),
		       COALESCE(activo, true), COALESCE(fecha_modificacion::text, ''), 
		       COALESCE(fecha_creacion::text, '')
		FROM blog."articulos_secciones"
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var secciones []models.ArticuloSeccion
	for rows.Next() {
		var s models.ArticuloSeccion
		err := rows.Scan(
			&s.ID, &s.ArticuloID, &s.TituloSeccion, &s.Contenido,
			&s.ImagenURL, &s.Orden, &s.Activo, &s.FechaModificacion, &s.FechaCreacion,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		secciones = append(secciones, s)
	}
	c.JSON(http.StatusOK, secciones)
}