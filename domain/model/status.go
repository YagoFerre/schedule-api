package model

type Status string

const (
	Agendado  Status = "agendado"
	Enviado   Status = "enviado"
	Cancelado Status = "cancelado"
)
