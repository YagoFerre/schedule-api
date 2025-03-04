package model

type StatusModel string

const (
	Agendado  StatusModel = "agendado"
	Enviado   StatusModel = "enviado"
	Cancelado StatusModel = "cancelado"
)
