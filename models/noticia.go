package models

import "database/sql"

type Noticia struct {
	ID              int            `json:"id"`
	CategoriaID     int            `json:"categoria_id"`
	Titulo          string         `json:"titulo"`
	Contenido       string         `json:"contenido"`
	Encabezado      sql.NullString `json:"encabezado"`
	ImagenPrincipal sql.NullString `json:"imagen_principal"`
	AutorID         int            `json:"autor_id"`
	Estado          string         `json:"estado"`
	Vistas          int            `json:"vistas"`
	PublicadoEn     sql.NullString `json:"publicado_en"`
	CreadoEn        string         `json:"creado_en"`
	ActualizadoEn   string         `json:"actualizado_en"`
	Activo          bool           `json:"activo"`
}