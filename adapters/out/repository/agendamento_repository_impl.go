package repository

import (
	"schedule-api/application/domain/model"
	"schedule-api/application/domain/port/out/repository"
	"schedule-api/configuration/rest_errors"

	"gorm.io/gorm"
)

func NewAgendamentoRepository(database *gorm.DB) repository.AgendamentoRepository {
	return &agendamentoRepository{
		database,
	}
}

type agendamentoRepository struct {
	database *gorm.DB
}

func (r *agendamentoRepository) Create(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr) {
	result := r.database.Create(&agendamentoModel)
	if result.Error != nil {
		return &model.AgendamentoModel{}, &rest_errors.RestErr{
			Message: "Falha ao salvar agendamento",
			Code:    500,
		}
	}

	return &agendamentoModel, nil
}

func (r *agendamentoRepository) FindById(id int) (*model.AgendamentoModel, *rest_errors.RestErr) {
	agendamentoModel := &model.AgendamentoModel{}

	result := r.database.First(&agendamentoModel, id)
	if result.Error != nil {
		return nil, rest_errors.NewNotFoundError("Agendamento não encontrado")
	}

	return agendamentoModel, nil
}

func (r *agendamentoRepository) Update(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
