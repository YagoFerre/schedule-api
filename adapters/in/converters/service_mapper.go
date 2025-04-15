package converters

import (
	"schedule-api/adapters/out/domain/entities"
	"schedule-api/adapters/out/domain/request"
	"schedule-api/adapters/out/domain/response"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoMapper interface {
	HandleMapperAgendarEnvioNotificacao(agendamentoRequest *request.AgendamentoRequest) (*entities.AgendamentoEntity, *rest_errors.RestErr)
	HandleMapperConsultarEnvioNotificacao(id uint) (*response.AgendamentoResponse, *rest_errors.RestErr)
	HandleMapperCancelarEnvioNotificacao(id uint) (*entities.AgendamentoEntity, *rest_errors.RestErr)
}
