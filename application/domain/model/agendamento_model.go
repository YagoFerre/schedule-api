package model

import "time"

type AgendamentoModel struct {
	ID                  uint
	Destinatario        string
	DataHoraEnvio       time.Time
	DataHoraAgendamento time.Time
	DataHoraModificacao time.Time
	Mensagem            string
	StatusNotificacao   StatusModel
}

func (AgendamentoModel) TableName() string {
	return "agendamento_entities"
}
