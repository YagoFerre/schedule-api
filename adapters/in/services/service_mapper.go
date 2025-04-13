package mapper

import (
	"schedule-api/application/domain/model"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoMapper interface {
	HandleMapperAgendarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
	HandleMapperConsultarEnvioNotificacao(id int) (*model.AgendamentoModel, *rest_errors.RestErr)
	HandleMapperCancelarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
}
