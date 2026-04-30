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

// CreateComentarioComunidad crea un nuevo comentario
func CreateComentarioComunidad(c *gin.Context) {
	var coment models.ComentarioComunidad
	if err := c.ShouldBindJSON(&coment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog."comentarios_comunidad" (categoria_id, usuario_id, contenido, 
		       calificacion, likes, dislikes, estado, activo)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, coment.CategoriaID, coment.UsuarioID, coment.Contenido,
		coment.Calificacion, coment.Likes, coment.Dislikes, coment.Estado, coment.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	coment.ID = id
	c.JSON(http.StatusCreated, coment)
}

// UpdateComentarioComunidad actualiza un comentario existente
func UpdateComentarioComunidad(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var coment models.ComentarioComunidad
	if err := c.ShouldBindJSON(&coment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE blog."comentarios_comunidad" 
		SET categoria_id = $1, usuario_id = $2, contenido = $3, 
		    calificacion = $4, likes = $5, dislikes = $6, 
		    estado = $7, activo = $8
		WHERE id = $9
	`
	result, err := config.DB.Exec(query, coment.CategoriaID, coment.UsuarioID, coment.Contenido,
		coment.Calificacion, coment.Likes, coment.Dislikes, coment.Estado, coment.Activo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		return
	}
	coment.ID = id
	c.JSON(http.StatusOK, coment)
}

// DeleteComentarioComunidad elimina un comentario
func DeleteComentarioComunidad(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM blog.\"comentarios_comunidad\" WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}