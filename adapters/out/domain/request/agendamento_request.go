package request

import "time"

type AgendamentoRequest struct {
	DataHoraEnvio time.Time `json:"dataHoraEnvio" binding:"required"`
	Destinatario  string    `json:"destinatario" binding:"required"`
	Mensagem      string    `json:"mensagem" binding:"required,min=4,max=200"`
}
