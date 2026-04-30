package models

type ComentarioComunidad struct {
	ID                int    `json:"id"`
	CategoriaID       int    `json:"categoria_id"`
	UsuarioID         int    `json:"usuario_id"`
	Contenido         string `json:"contenido"`
	Calificacion      int    `json:"calificacion"`
	Likes             int    `json:"likes"`
	Dislikes          int    `json:"dislikes"`
	Estado            string `json:"estado"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}


