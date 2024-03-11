package services

import (
	"avitoCase/pkg/repository"
)

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService {
	return &OperationService{repo: repo}
}

func (s *OperationService) GetData() {
	//	TODO Логика
}
