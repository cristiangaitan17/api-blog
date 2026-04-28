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

func CreateNoticia(c *gin.Context) {
	var n models.Noticia
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO blog.noticias (categoria_id, titulo, contenido, encabezado, 
		       imagen_principal, autor_id, estado, vistas, publicado_en,
		       creado_en, actualizado_en, activo)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW(), $10)
		RETURNING id, creado_en, actualizado_en
	`
	err := config.DB.QueryRow(query, n.CategoriaID, n.Titulo, n.Contenido, n.Encabezado,
		n.ImagenPrincipal, n.AutorID, n.Estado, n.Vistas, n.PublicadoEn, n.Activo).
		Scan(&n.ID, &n.CreadoEn, &n.ActualizadoEn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, n)
}

func UpdateNoticia(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var n models.Noticia
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE blog.noticias 
		SET categoria_id = $1, titulo = $2, contenido = $3, encabezado = $4,
		    imagen_principal = $5, autor_id = $6, estado = $7, vistas = $8, 
		    publicado_en = $9, actualizado_en = NOW(), activo = $10
		WHERE id = $11
		RETURNING id, categoria_id, titulo, contenido, encabezado, imagen_principal,
		          autor_id, estado, vistas, publicado_en, creado_en, actualizado_en, activo
	`
	row := config.DB.QueryRow(query, n.CategoriaID, n.Titulo, n.Contenido, n.Encabezado,
		n.ImagenPrincipal, n.AutorID, n.Estado, n.Vistas, n.PublicadoEn, n.Activo, id)
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

func DeleteNoticia(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM blog.noticias WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Noticia no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}