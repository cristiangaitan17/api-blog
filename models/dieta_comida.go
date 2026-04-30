package models

type DietaComida struct {
	ID                int    `json:"id"`
	DietaID           int    `json:"dieta_id"`
	TiempoComida      string `json:"tiempo_comida"`
	Descripcion       string `json:"descripcion"`
	Orden             int    `json:"orden"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}

