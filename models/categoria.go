package models

type Categoria struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	SeccionLugar      string `json:"seccion_lugar"`
	Descripcion       string `json:"descripcion"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}

 