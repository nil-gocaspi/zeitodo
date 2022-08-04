package service

import (
	"example-project/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	//	UpdateMany(docs []interface{}) interface{}
	GetByID(id string) model.Employee
	GetProposals(id string) ([]model.Proposal, error)
	SaveProposals(docs []interface{}) (interface{}, error)
}

type EmployeeService struct {
	DbService DatabaseInterface
}

func NewEmployeeService(dbInterface DatabaseInterface) EmployeeService {
	return EmployeeService{
		DbService: dbInterface,
	}
}

func (s EmployeeService) GetEmployeeById(id string) model.Employee {
	return s.DbService.GetByID(id)
}
