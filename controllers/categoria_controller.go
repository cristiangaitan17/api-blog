package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetCategorias obtiene todas las categorías
func GetCategorias(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, nombre, seccion_lugar, descripcion, activo, 
		       Fecha_modificacion, Fecha_creacion 
		FROM blog.categorias
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var cat models.Categoria
		err := rows.Scan(
			&cat.ID, &cat.Nombre, &cat.SeccionLugar, &cat.Descripcion,
			&cat.Activo, &cat.FechaModificacion, &cat.FechaCreacion,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categorias = append(categorias, cat)
	}
	c.JSON(http.StatusOK, categorias)
}

// GetCategoriaByID obtiene una categoría por ID
func GetCategoriaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cat models.Categoria
	row := config.DB.QueryRow(`
		SELECT id, nombre, seccion_lugar, descripcion, activo, 
		       Fecha_modificacion, Fecha_creacion 
		FROM blog.categorias WHERE id = $1
	`, id)

	err = row.Scan(
		&cat.ID, &cat.Nombre, &cat.SeccionLugar, &cat.Descripcion,
		&cat.Activo, &cat.FechaModificacion, &cat.FechaCreacion,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// CreateCategoria crea una nueva categoría
func CreateCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog.categorias (nombre, seccion_lugar, descripcion, activo, Fecha_modificacion, Fecha_creacion)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, Fecha_modificacion, Fecha_creacion
	`
	err := config.DB.QueryRow(query, cat.Nombre, cat.SeccionLugar, cat.Descripcion, cat.Activo).
		Scan(&cat.ID, &cat.FechaModificacion, &cat.FechaCreacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// UpdateCategoria actualiza una categoría existente
func UpdateCategoria(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cat models.Categoria
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE blog.categorias 
		SET nombre = $1, seccion_lugar = $2, descripcion = $3, activo = $4, Fecha_modificacion = NOW()
		WHERE id = $5
		RETURNING id, nombre, seccion_lugar, descripcion, activo, Fecha_modificacion, Fecha_creacion
	`
	row := config.DB.QueryRow(query, cat.Nombre, cat.SeccionLugar, cat.Descripcion, cat.Activo, id)
	err = row.Scan(
		&cat.ID, &cat.Nombre, &cat.SeccionLugar, &cat.Descripcion,
		&cat.Activo, &cat.FechaModificacion, &cat.FechaCreacion,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// DeleteCategoria elimina una categoría
func DeleteCategoria(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM blog.categorias WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}