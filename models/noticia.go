package models

import "database/sql"

type Noticia struct {
	ID              int            `json:"id"`
	CategoriaID     sql.NullInt64  `json:"categoria_id"`
	Titulo          string         `json:"titulo"`
	Contenido       string         `json:"contenido"`
	Encabezado      sql.NullString `json:"encabezado"`
	ImagenPrincipal sql.NullString `json:"imagen_principal"`
	AutorID         sql.NullInt64  `json:"autor_id"`
	Estado          string         `json:"estado"`
	Vistas          int            `json:"vistas"`
	PublicadoEn     sql.NullString `json:"publicado_en"`
	CreadoEn        sql.NullString `json:"creado_en"`      // Puede ser NULL
	ActualizadoEn   sql.NullString `json:"actualizado_en"` // Puede ser NULL
	Activo          bool           `json:"activo"`
}