package domain

import (
	"database/sql"

	"github.com/SaravananPitchaimuthu/Fruits/Fruits/logger"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/utils/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type FruitsRepositoryDb struct {
	client *sqlx.DB
}

func (d FruitsRepositoryDb) ById(id string) (*Fruit, *errors.AppError) {
	findByIdQuery := "SELECT fruit_id,name,price,quantity FROM fruits WHERE fruit_id = ? "
	// row := d.client.QueryRow(findByIdQuery, id)

	var f Fruit
	err := d.client.Get(&f, findByIdQuery, id)
	// err := row.Scan(&f.Id, &f.Name, &f.Price, &f.Quantity)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Fruit Not Found")
			return nil, errors.NewNotFoundError("Fruit not Found")
		} else {
			logger.Error("Unexpected database error")
			return nil, errors.NewInternalServerError("Unexpected database error")
		}
	}

	return &f, nil

}

func (f FruitsRepositoryDb) FindAll(status string) ([]Fruit, *errors.AppError) {
	fruits := make([]Fruit, 0)
	var err error
	if status == "" {
		findAllQuery := "SELECT fruit_id,name,price,quantity FROM fruits"
		err = f.client.Select(&fruits, findAllQuery)
	} else {
		findAllQuery := "SELECT fruit_id,name,price,quantity FROM fruits where status = ?"
		err = f.client.Select(&fruits, findAllQuery, status)
	}

	if err != nil {
		logger.Error("Error while scanning fruits")
		return nil, errors.NewInternalServerError("Unexpected database error")
	}

	// rows, err := f.client.Query(findAllQuery)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	// fruits := make([]Fruit, 0)
	// err = sqlx.StructScan(rows, &fruits)

	// if err != nil {
	// 	logger.Error("Error while scanning fruits")
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	var f Fruit
	// 	if err := rows.Scan(&f.Id, &f.Name, &f.Price, &f.Quantity); err != nil {
	// 		logger.Error("Error while scanning fruits")
	// 		return nil, err
	// 	}
	// 	fruits = append(fruits, f)
	// }

	return fruits, nil
}

func NewFruitRepositoryDb(client *sqlx.DB) FruitsRepositoryDb {
	return FruitsRepositoryDb{client}
}
