package dto

type FruitResponse struct {
	Id       string `json:"fruit_id"`
	Name     string `json:"full_name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
