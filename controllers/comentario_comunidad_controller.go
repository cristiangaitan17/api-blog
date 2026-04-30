package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

// GetComentariosComunidad obtiene todos los comentarios
func GetComentariosComunidad(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, COALESCE(categoria_id, 0), COALESCE(usuario_id, 0), 
		       COALESCE(contenido, ''), COALESCE(calificacion, 0), 
		       COALESCE(likes, 0), COALESCE(dislikes, 0), 
		       COALESCE(estado, 'activo'), COALESCE(activo, true)
		FROM blog."comentarios_comunidad"
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var comentarios []models.ComentarioComunidad
	for rows.Next() {
		var coment models.ComentarioComunidad
		err := rows.Scan(
			&coment.ID, &coment.CategoriaID, &coment.UsuarioID, &coment.Contenido,
			&coment.Calificacion, &coment.Likes, &coment.Dislikes,
			&coment.Estado, &coment.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comentarios = append(comentarios, coment)
	}
	c.JSON(http.StatusOK, comentarios)
}
// GetComentarioComunidadByID obtiene un comentario por ID
func GetComentarioComunidadByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var coment models.ComentarioComunidad
	row := config.DB.QueryRow(`
		SELECT id, COALESCE(categoria_id, 0), COALESCE(usuario_id, 0), 
		       COALESCE(contenido, ''), COALESCE(calificacion, 0), 
		       COALESCE(likes, 0), COALESCE(dislikes, 0), 
		       COALESCE(estado, 'activo'), COALESCE(activo, true)
		FROM blog."comentarios_comunidad" WHERE id = $1
	`, id)

	err = row.Scan(
		&coment.ID, &coment.CategoriaID, &coment.UsuarioID, &coment.Contenido,
		&coment.Calificacion, &coment.Likes, &coment.Dislikes,
		&coment.Estado, &coment.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coment)
}
