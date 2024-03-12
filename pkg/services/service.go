package services

import (
	"avitoCase/pkg/repository"
)

type CostMatrix interface {
	GetData(userId int, locationId int, microcategoryId int) (map[string]string, error)
}

type Service struct {
	CostMatrix
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CostMatrix: NewCostMatrixService(repos.CostMatrix),
	}
}
