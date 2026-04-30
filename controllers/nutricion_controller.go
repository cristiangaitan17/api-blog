package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetNutricion obtiene todos los planes de nutrición
func GetNutricion(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, nombre, descripcion, objetivo, 
		       COALESCE(imagen_url, ''), 
		       COALESCE(autor_id, 0), 
		       COALESCE(publicado, false), 
		       COALESCE(creado_en::text, ''), 
		       COALESCE(activo, true)
		FROM blog."Nutricion"
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var nutricion []models.Nutricion
	for rows.Next() {
		var n models.Nutricion
		err := rows.Scan(
			&n.ID, &n.Nombre, &n.Descripcion, &n.Objetivo,
			&n.ImagenURL, &n.AutorID, &n.Publicado,
			&n.CreadoEn, &n.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		nutricion = append(nutricion, n)
	}
	c.JSON(http.StatusOK, nutricion)
}

// GetNutricionByID obtiene un plan de nutrición por ID
func GetNutricionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var n models.Nutricion
	row := config.DB.QueryRow(`
		SELECT id, nombre, descripcion, objetivo, 
		       COALESCE(imagen_url, ''), 
		       COALESCE(autor_id, 0), 
		       COALESCE(publicado, false), 
		       COALESCE(creado_en::text, ''), 
		       COALESCE(activo, true)
		FROM blog."Nutricion" WHERE id = $1
	`, id)

	err = row.Scan(
		&n.ID, &n.Nombre, &n.Descripcion, &n.Objetivo,
		&n.ImagenURL, &n.AutorID, &n.Publicado,
		&n.CreadoEn, &n.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan de nutrición no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, n)
}

// CreateNutricion crea un nuevo plan de nutrición
func CreateNutricion(c *gin.Context) {
	var n models.Nutricion
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog."Nutricion" (nombre, descripcion, objetivo, imagen_url, autor_id, 
		       publicado, creado_en, activo)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, n.Nombre, n.Descripcion, n.Objetivo, n.ImagenURL,
		n.AutorID, n.Publicado, n.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	n.ID = id
	c.JSON(http.StatusCreated, n)
}
