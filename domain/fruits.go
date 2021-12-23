package domain

import (
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/dto"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/utils/errors"
)

type Fruit struct {
	Id       string `db:"fruit_id"`
	Name     string
	Price    string
	Quantity string
	Status   string
}

func (f Fruit) ToDto() dto.FruitResponse {
	return dto.FruitResponse{
		Id:       f.Id,
		Name:     f.Name,
		Price:    f.Price,
		Quantity: f.Quantity,
	}
}

type FruitRepository interface {
	FindAll(status string) ([]Fruit, *errors.AppError)
	ById(string) (*Fruit, *errors.AppError)
}
