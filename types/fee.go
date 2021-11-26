package types

import "time"

type FeeResponse struct {
	ExpirationDate time.Time `json:"expirationDate"`
	Amount         string    `json:"amount"`
}

type FeeParams struct {
	SellToken string `json:"sellToken"`
	BuyToken  string `json:"buyToken"`
	Amount    string `json:"amount"`
	Kind      string `json:"kind"`
}
