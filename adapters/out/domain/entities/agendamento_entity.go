package entities

import (
	"time"
)

type AgendamentoEntity struct {
	ID                  uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Destinatario        string    `json:"destinatario"`
	DataHoraEnvio       time.Time `json:"dataHoraEnvio"`
	DataHoraAgendamento time.Time `json:"dataHoraAgendamento" gorm:"autoCreateTime"`
	DataHoraModificacao time.Time `json:"dataHoraModificacao" gorm:"autoUpdateTime"`
	Mensagem            string    `json:"mensagem"`
	StatusNotificacao   Status    `json:"statusNotificacao" gorm:"type:status"`
}
