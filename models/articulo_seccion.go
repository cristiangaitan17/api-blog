package models

type ArticuloSeccion struct {
	ID                int    `json:"id"`
	ArticuloID        int    `json:"articulo_id"`
	TituloSeccion     string `json:"titulo_seccion"`
	Contenido         string `json:"contenido"`
	ImagenURL         string `json:"imagen_url"`
	Orden             int    `json:"orden"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}