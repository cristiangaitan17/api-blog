package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api-blog/config"
	"github.com/cristiangaitan17/api-blog/models"
)

func GetNoticias(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, categoria_id, titulo, contenido, encabezado, 
		       imagen_principal, autor_id, estado, vistas, publicado_en,
		       creado_en, actualizado_en, activo
		FROM blog.noticias
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var noticias []models.Noticia
	for rows.Next() {
		var n models.Noticia
		err := rows.Scan(
			&n.ID, &n.CategoriaID, &n.Titulo, &n.Contenido, &n.Encabezado,
			&n.ImagenPrincipal, &n.AutorID, &n.Estado, &n.Vistas, &n.PublicadoEn,
			&n.CreadoEn, &n.ActualizadoEn, &n.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		noticias = append(noticias, n)
	}
	c.JSON(http.StatusOK, noticias)
}

func GetNoticiaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var n models.Noticia
	row := config.DB.QueryRow(`
		SELECT id, categoria_id, titulo, contenido, encabezado, 
		       imagen_principal, autor_id, estado, vistas, publicado_en,
		       creado_en, actualizado_en, activo
		FROM blog.noticias WHERE id = $1
	`, id)

	err = row.Scan(
		&n.ID, &n.CategoriaID, &n.Titulo, &n.Contenido, &n.Encabezado,
		&n.ImagenPrincipal, &n.AutorID, &n.Estado, &n.Vistas, &n.PublicadoEn,
		&n.CreadoEn, &n.ActualizadoEn, &n.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Noticia no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, n)
}