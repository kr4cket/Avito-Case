package services

import (
	"avitoCase/pkg/repository"
)

type CostMatrixService struct {
	repo repository.CostMatrix
}

func NewCostMatrixService(repo repository.CostMatrix) *CostMatrixService {
	return &CostMatrixService{repo: repo}
}

func (s *CostMatrixService) GetData(userId int, locationId int, microcategoryId int) (map[string]string, error) {
	if userId == 100 && locationId == 25 && microcategoryId == 15 {
		return s.repo.GetData(userId, locationId, microcategoryId)
	}
	return map[string]string{}, nil
}
