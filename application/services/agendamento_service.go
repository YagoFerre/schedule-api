package services

import (
	"schedule-api/application/domain/model"
	"schedule-api/application/domain/port/in/usecases"
	"schedule-api/application/domain/port/out/repository"
	"schedule-api/configuration/rest_errors"
)

func NewAgandamentoService(agendamentoRepository repository.AgendamentoRepository) usecases.AgendamentoUseCases {
	return &agendamentoService{
		agendamentoRepository,
	}
}

type agendamentoService struct {
	agendamentoRepository repository.AgendamentoRepository
}

func (au *agendamentoService) ExecuteAgendarEnvioNotificacao(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr) {
	return au.agendamentoRepository.AgendarEnvioNotificacao(agendamentoModel)
}

func (au *agendamentoService) ExecuteConsultarEnvioNotificacao(id int) (*model.AgendamentoModel, *rest_errors.RestErr) {
	return au.agendamentoRepository.ConsultarEnvioNotificacao(id)
}

func (au *agendamentoService) ExecuteCancelarEnvioNotificacao(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr) {
	return au.agendamentoRepository.CancelarEnvioNotificacao(agendamentoModel)
}
