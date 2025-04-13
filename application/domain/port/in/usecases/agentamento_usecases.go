package usecases

import (
	"schedule-api/application/domain/model"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoUseCases interface {
	ExecuteAgendarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
	ExecuteConsultarEnvioNotificacao(id uint) (*model.AgendamentoModel, *rest_errors.RestErr)
	ExecuteCancelarEnvioNotificacao(id uint) (*model.AgendamentoModel, *rest_errors.RestErr)
}
