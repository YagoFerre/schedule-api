package mapper

import (
	"schedule-api/adapters/out/domain/entities"
	"schedule-api/adapters/out/domain/request"
	"schedule-api/adapters/out/domain/response"
	"schedule-api/application/domain/model"
)

func ConvertRequestToModel(agendamentoRequest *request.AgendamentoRequest) model.AgendamentoModel {
	return model.AgendamentoModel{
		DataHoraEnvio: agendamentoRequest.DataHoraEnvio,
		Destinatario:  agendamentoRequest.Destinatario,
		Mensagem:      agendamentoRequest.Mensagem,
	}
}

func ConvertModelToEntity(agendamentoModel *model.AgendamentoModel) entities.AgendamentoEntity {
	return entities.AgendamentoEntity{
		ID:                  agendamentoModel.ID,
		Destinatario:        agendamentoModel.Destinatario,
		DataHoraEnvio:       agendamentoModel.DataHoraEnvio,
		DataHoraAgendamento: agendamentoModel.DataHoraAgendamento,
		DataHoraModificacao: agendamentoModel.DataHoraModificacao,
		Mensagem:            agendamentoModel.Mensagem,
		StatusNotificacao:   entities.Status(agendamentoModel.StatusNotificacao),
	}
}

func ConvertModelToResponse(agendamentoModel *model.AgendamentoModel) response.AgendamentoResponse {
	return response.AgendamentoResponse{
		ID:                agendamentoModel.ID,
		Destinatario:      agendamentoModel.Destinatario,
		DataHoraEnvio:     agendamentoModel.DataHoraEnvio,
		Mensagem:          agendamentoModel.Mensagem,
		StatusNotificacao: entities.Status(agendamentoModel.StatusNotificacao),
	}
}
