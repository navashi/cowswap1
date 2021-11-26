package types

type Order struct {
	SellToken                    string `json:"sellToken"`
	BuyToken                     string `json:"buyToken"`
	Receiver                     string `json:"receiver"`
	SellAmount                   string `json:"sellAmount"`
	BuyAmount                    string `json:"buyAmount"`
	ValidTo                      int64  `json:"validTo"`
	AppData                      string `json:"appData"`
	FeeAmount                    string `json:"feeAmount"`
	Kind                         string `json:"kind"`
	PartiallyFillable            bool   `json:"partiallyFillable"`
	SellTokenBalance             string `json:"sellTokenBalance"`
	BuyTokenBalance              string `json:"buyTokenBalance"`
	SigningScheme                string `json:"signingScheme"`
	Signature                    string `json:"signature"`
	From                         string `json:"from"`
	CreationTime                 string `json:"creationTime"`
	Owner                        string `json:"owner"`
	UID                          string `json:"uid"`
	AvailableBalance             string `json:"availableBalance"`
	ExecutedSellAmount           string `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string `json:"executedSellAmountBeforeFees"`
	ExecutedBuyAmount            string `json:"executedBuyAmount"`
	ExecutedFeeAmount            string `json:"executedFeeAmount"`
	Invalidated                  bool   `json:"invalidated"`
	Status                       string `json:"status"`
}

type OrdersParams struct {
	Owner                      string `json:"owner"`
	SellToken                  string `json:"sellToken"`
	BuyToken                   string `json:"buyToken"`
	IncludeFullyExecuted       bool   `json:"includeFullyExecuted"`
	IncludeInvalidated         bool   `json:"includeInvalidated"`
	IncludeInsufficientBalance bool   `json:"includeInsufficientBalance"`
	IncludePresignaturePending bool   `json:"includePresignaturePending"`
	IncludeUnsupportedTokens   bool   `json:"includeUnsupportedTokens"`
	MinValidTo                 int64  `json:"minValidTo"`
}

type OrderParams struct {
	SellToken         string  `json:"sellToken"`
	BuyToken          string  `json:"buyToken"`
	SellAmount        string  `json:"sellAmount"`
	BuyAmount         string  `json:"buyAmount"`
	ValidTo           int64   `json:"validTo"`
	AppData           string  `json:"appData"`
	FeeAmount         string  `json:"feeAmount"`
	Kind              string  `json:"kind"`
	Receiver          *string `json:"receiver,omitempty"`
	PartiallyFillable bool    `json:"partiallyFillable"`
	SellTokenBalance  *string `json:"sellTokenBalance,omitempty"`
	BuyTokenBalance   *string `json:"buyTokenBalance,omitempty"`
	SigningScheme     string  `json:"signingScheme"`
	Signature         string  `json:"signature"`
	From              *string `json:"from,omitempty"`
}
