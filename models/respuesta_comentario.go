package models

type RespuestaComentario struct {
	ID                int    `json:"id"`
	ComentarioID      int    `json:"comentario_id"`
	UsuarioID         int    `json:"usuario_id"`
	Contenido         string `json:"contenido"`
	CreadoEn          string `json:"creado_en"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}