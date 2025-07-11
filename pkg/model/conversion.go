package model

type ConversionRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	From   string  `json:"from" validate:"required,len=3"`
	To     string  `json:"to" validate:"required,len=3"`
}

type ConversionResponse struct {
	ConvertedAmount float64 `json:"converted_amount"`
	Rate            float64 `json:"rate"`
}
