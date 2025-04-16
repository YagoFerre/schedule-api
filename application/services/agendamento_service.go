package services

import (
	"schedule-api/application/domain/model"
	"schedule-api/application/domain/port/in/usecases"
	"schedule-api/application/domain/port/out/repository"
	"schedule-api/configuration/rest_errors"
	"time"
)

// Impl Service usecases
func NewAgandamentoService(agendamentoRepository repository.AgendamentoRepository) usecases.AgendamentoUseCases {
	return &agendamentoService{
		agendamentoRepository,
	}
}

type agendamentoService struct {
	agendamentoRepository repository.AgendamentoRepository
}

func (as *agendamentoService) ExecuteAgendarEnvioNotificacao(agendamentoModel *model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr) {
	agendamentoModel.StatusNotificacao = model.Agendado
	agendamentoModel.DataHoraModificacao = time.Now()
	agendamentoModel.DataHoraAgendamento = time.Now()
	return as.agendamentoRepository.Create(agendamentoModel)
}

func (as *agendamentoService) ExecuteConsultarEnvioNotificacao(id uint) (*model.AgendamentoModel, *rest_errors.RestErr) {
	return as.agendamentoRepository.FindById(id)
}

func (as *agendamentoService) ExecuteCancelarEnvioNotificacao(id uint) (*model.AgendamentoModel, *rest_errors.RestErr) {
	agendamentoModel := &model.AgendamentoModel{
		ID:                  id,
		StatusNotificacao:   model.Cancelado,
		DataHoraModificacao: time.Now(),
	}

	return as.agendamentoRepository.Update(agendamentoModel)
}
