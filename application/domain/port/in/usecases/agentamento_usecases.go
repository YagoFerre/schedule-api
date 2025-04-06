package usecases

import (
	"schedule-api/application/domain/model"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoUseCases interface {
	ExecuteAgendarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
	ExecuteConsultarEnvioNotificacao(id int) (*model.AgendamentoModel, *rest_errors.RestErr)
	ExecuteCancelarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
}
