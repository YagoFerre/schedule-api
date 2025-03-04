package repository

import (
	"schedule-api/application/domain/model"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoRepository interface {
	AgendarEnvioNotificacao(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
	ConsultarEnvioNotificacao(id int) (*model.AgendamentoModel, *rest_errors.RestErr)
	CancelarEnvioNotificacao(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
}
