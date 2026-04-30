package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetRespuestasComentario obtiene todas las respuestas
func GetRespuestasComentario(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, COALESCE(comentario_id, 0), COALESCE(usuario_id, 0), 
		       COALESCE(contenido, ''), COALESCE(creado_en::text, ''), 
		       COALESCE(activo, true)
		FROM blog."respuestas_comentario"
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var respuestas []models.RespuestaComentario
	for rows.Next() {
		var r models.RespuestaComentario
		err := rows.Scan(
			&r.ID, &r.ComentarioID, &r.UsuarioID, &r.Contenido,
			&r.CreadoEn, &r.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		respuestas = append(respuestas, r)
	}
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestaComentarioByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var r models.RespuestaComentario
	row := config.DB.QueryRow(`
		SELECT id, COALESCE(comentario_id, 0), COALESCE(usuario_id, 0), 
		       COALESCE(contenido, ''), COALESCE(creado_en::text, ''), 
		       COALESCE(activo, true)
		FROM blog."respuestas_comentario" WHERE id = $1
	`, id)

	err = row.Scan(
		&r.ID, &r.ComentarioID, &r.UsuarioID, &r.Contenido,
		&r.CreadoEn, &r.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Respuesta no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// CreateRespuestaComentario crea una nueva respuesta
func CreateRespuestaComentario(c *gin.Context) {
	var r models.RespuestaComentario
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog."respuestas_comentario" (comentario_id, usuario_id, contenido, creado_en, activo)
		VALUES ($1, $2, $3, NOW(), $4)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, r.ComentarioID, r.UsuarioID, r.Contenido, r.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r.ID = id
	c.JSON(http.StatusCreated, r)
}