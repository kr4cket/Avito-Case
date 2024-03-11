package services

import (
	"avitoCase/pkg/repository"
)

type Operation interface {
	GetData()
}

type Service struct {
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{

		Operation: NewOperationService(repos.Operation),
	}
}
