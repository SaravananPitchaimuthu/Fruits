package service

import (
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/domain"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/dto"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/utils/errors"
)

type FruitService interface {
	GetAllFruits(string) ([]domain.Fruit, *errors.AppError)
	GetFruit(string) (*dto.FruitResponse, *errors.AppError)
}

type DefaultFruitService struct {
	repo domain.FruitRepository
}

func (f DefaultFruitService) GetAllFruits(status string) ([]domain.Fruit, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return f.repo.FindAll(status)
}

func (f DefaultFruitService) GetFruit(id string) (*dto.FruitResponse, *errors.AppError) {
	resp, err := f.repo.ById(id)

	if err != nil {
		return nil, err
	}
	response := resp.ToDto()
	return &response, nil
}

func NewFruitService(repository domain.FruitRepository) DefaultFruitService {
	return DefaultFruitService{repository}
}
