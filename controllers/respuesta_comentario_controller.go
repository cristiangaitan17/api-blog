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