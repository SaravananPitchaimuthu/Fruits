package dto

import (
	"strings"

	"github.com/SaravananPitchaimuthu/Fruits/Fruits/utils/errors"
)

type NewAccountRequest struct {
	FruitId     string  `json:"fruit_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidationError("to open new account you need to deposit atleast 5000")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidationError("account type should be saving or checking")
	}
	return nil
}