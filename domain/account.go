package domain

import (
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/dto"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/utils/errors"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	FruitId     string  `db:"fruit_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
	FindBy(string) (*Account, *errors.AppError)
	SaveTransaction(Transaction) (*Transaction, *errors.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
