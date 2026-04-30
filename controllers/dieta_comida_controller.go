package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetDietaComidas obtiene todas las comidas de dieta
func GetDietaComidas(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, COALESCE(dieta_id, 0), COALESCE(tiempo_comida, ''), 
		       COALESCE(descripcion, ''), COALESCE(orden, 0), 
		       COALESCE(activo, true)
		FROM blog."dieta_comidas"
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var comidas []models.DietaComida
	for rows.Next() {
		var d models.DietaComida
		err := rows.Scan(
			&d.ID, &d.DietaID, &d.TiempoComida, &d.Descripcion,
			&d.Orden, &d.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comidas = append(comidas, d)
	}
	c.JSON(http.StatusOK, comidas)
}
// GetDietaComidaByID obtiene una comida por ID
func GetDietaComidaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var d models.DietaComida
	row := config.DB.QueryRow(`
		SELECT id, COALESCE(dieta_id, 0), COALESCE(tiempo_comida, ''), 
		       COALESCE(descripcion, ''), COALESCE(orden, 0), 
		       COALESCE(activo, true)
		FROM blog."dieta_comidas" WHERE id = $1
	`, id)

	err = row.Scan(
		&d.ID, &d.DietaID, &d.TiempoComida, &d.Descripcion,
		&d.Orden, &d.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comida no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, d)
}

// CreateDietaComida crea una nueva comida
func CreateDietaComida(c *gin.Context) {
	var d models.DietaComida
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog."dieta_comidas" (dieta_id, tiempo_comida, descripcion, orden, activo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, d.DietaID, d.TiempoComida, d.Descripcion,
		d.Orden, d.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	d.ID = id
	c.JSON(http.StatusCreated, d)
}

// UpdateDietaComida actualiza una comida existente
func UpdateDietaComida(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var d models.DietaComida
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE blog."dieta_comidas" 
		SET dieta_id = $1, tiempo_comida = $2, descripcion = $3, orden = $4, activo = $5
		WHERE id = $6
	`
	result, err := config.DB.Exec(query, d.DietaID, d.TiempoComida, d.Descripcion,
		d.Orden, d.Activo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comida no encontrada"})
		return
	}
	d.ID = id
	c.JSON(http.StatusOK, d)
}

// DeleteDietaComida elimina una comida
func DeleteDietaComida(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM blog.\"dieta_comidas\" WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comida no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}