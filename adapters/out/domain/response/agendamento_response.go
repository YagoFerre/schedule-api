package response

import (
	"schedule-api/adapters/out/domain/entities"
	"time"
)

type AgendamentoResponse struct {
	ID                uint            `json:"id"`
	Destinatario      string          `json:"destinatario"`
	DataHoraEnvio     time.Time       `json:"dataHoraEnvio"`
	Mensagem          string          `json:"mensagem"`
	StatusNotificacao entities.Status `json:"statusNotificacao"`
}
