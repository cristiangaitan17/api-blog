package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
)

type CategoriaSimple struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	SeccionLugar string `json:"seccion_lugar"`
	Descripcion string `json:"descripcion"`
	Activo      bool   `json:"activo"`
}

// GetCategorias obtiene todas las categorías
func GetCategorias(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, nombre, seccion_lugar, descripcion, activo FROM blog.categorias")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categorias []CategoriaSimple
	for rows.Next() {
		var cat CategoriaSimple
		if err := rows.Scan(&cat.ID, &cat.Nombre, &cat.SeccionLugar, &cat.Descripcion, &cat.Activo); err != nil {
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

	var cat CategoriaSimple
	row := config.DB.QueryRow("SELECT id, nombre, seccion_lugar, descripcion, activo FROM blog.categorias WHERE id = $1", id)
	if err := row.Scan(&cat.ID, &cat.Nombre, &cat.SeccionLugar, &cat.Descripcion, &cat.Activo); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// CreateCategoria crea una nueva categoría
func CreateCategoria(c *gin.Context) {
	var cat CategoriaSimple
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err := config.DB.QueryRow(
		"INSERT INTO blog.categorias (nombre, seccion_lugar, descripcion, activo) VALUES ($1, $2, $3, $4) RETURNING id",
		cat.Nombre, cat.SeccionLugar, cat.Descripcion, cat.Activo,
	).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cat.ID = id
	c.JSON(http.StatusCreated, cat)
}

// UpdateCategoria actualiza una categoría
func UpdateCategoria(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cat CategoriaSimple
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Exec(
		"UPDATE blog.categorias SET nombre = $1, seccion_lugar = $2, descripcion = $3, activo = $4 WHERE id = $5",
		cat.Nombre, cat.SeccionLugar, cat.Descripcion, cat.Activo, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	cat.ID = id
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
	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}