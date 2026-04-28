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

func GetArticuloSeccionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var s models.ArticuloSeccion
	row := config.DB.QueryRow(`
		SELECT id, COALESCE(articulo_id, 0), COALESCE(titulo_seccion, ''), 
		       COALESCE(contenido, ''), COALESCE(imagen_url, ''), COALESCE(orden, 0),
		       COALESCE(activo, true), COALESCE(fecha_modificacion::text, ''), 
		       COALESCE(fecha_creacion::text, '')
		FROM blog."articulos_secciones" WHERE id = $1
	`, id)

	err = row.Scan(
		&s.ID, &s.ArticuloID, &s.TituloSeccion, &s.Contenido,
		&s.ImagenURL, &s.Orden, &s.Activo, &s.FechaModificacion, &s.FechaCreacion,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sección no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func CreateArticuloSeccion(c *gin.Context) {
	var s models.ArticuloSeccion
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog."articulos_secciones" (articulo_id, titulo_seccion, contenido, imagen_url, 
		       orden, activo, fecha_modificacion, fecha_creacion)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, s.ArticuloID, s.TituloSeccion, s.Contenido,
		s.ImagenURL, s.Orden, s.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.ID = id
	c.JSON(http.StatusCreated, s)
}