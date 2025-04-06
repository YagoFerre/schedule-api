package repository

import (
	"schedule-api/application/domain/model"
	"schedule-api/configuration/rest_errors"
)

type AgendamentoRepository interface {
	Create(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
	FindById(id int) (*model.AgendamentoModel, *rest_errors.RestErr)
	Update(agendamentoModel model.AgendamentoModel) (*model.AgendamentoModel, *rest_errors.RestErr)
}
