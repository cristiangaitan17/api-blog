package models

type Nutricion struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	Objetivo          string `json:"objetivo"`
	ImagenURL         string `json:"imagen_url"`
	AutorID           int    `json:"autor_id"`
	Publicado         bool   `json:"publicado"`
	CreadoEn          string `json:"creado_en"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}
