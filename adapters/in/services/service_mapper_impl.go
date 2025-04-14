package services

import (
	"schedule-api/adapters/in/mapper"
	"schedule-api/adapters/out/domain/entities"
	"schedule-api/adapters/out/domain/request"
	"schedule-api/adapters/out/domain/response"
	"schedule-api/application/domain/port/in/usecases"
	"schedule-api/configuration/rest_errors"
)

func NewServiceMapperImpl(agendamentoUseCases usecases.AgendamentoUseCases) AgendamentoMapper {
	return &serviceMapper{
		agendamentoUseCases,
	}
}

type serviceMapper struct {
	agendamentoUseCases usecases.AgendamentoUseCases
}

func (s *serviceMapper) HandleMapperAgendarEnvioNotificacao(agendamentoRequest *request.AgendamentoRequest) (*entities.AgendamentoEntity, *rest_errors.RestErr) {
	agendamentoModel := mapper.ConvertRequestToModel(agendamentoRequest)

	result, err := s.agendamentoUseCases.ExecuteAgendarEnvioNotificacao(&agendamentoModel)
	if err != nil {
		return nil, err
	}

	agendamentoEntity := mapper.ConvertModelToEntity(result)

	return &agendamentoEntity, nil
}

func (s *serviceMapper) HandleMapperCancelarEnvioNotificacao(id uint) (*entities.AgendamentoEntity, *rest_errors.RestErr) {
	result, err := s.agendamentoUseCases.ExecuteCancelarEnvioNotificacao(id)
	if err != nil {
		return nil, err
	}

	agendamentoEntity := mapper.ConvertModelToEntity(result)

	return &agendamentoEntity, nil
}

func (s *serviceMapper) HandleMapperConsultarEnvioNotificacao(id uint) (*response.AgendamentoResponse, *rest_errors.RestErr) {
	result, err := s.agendamentoUseCases.ExecuteConsultarEnvioNotificacao(id)
	if err != nil {
		return nil, err
	}

	agendamentoResponse := mapper.ConvertModelToResponse(result)

	return &agendamentoResponse, nil
}
