package domain

type PaymentInfo struct {
	Transaction  string `json:"transaction" validate:"required,gte=10"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required,gte=2"`
	Provider     string `json:"provider" validate:"required,gte=4"`
	Amount       int64  `json:"amount" validate:"required,gte=0"`
	PaymentDT    int64  `json:"payment_dt" validate:"required,gte=0"`
	Bank         string `json:"bank" validate:"required,gte=4"`
	DeliveryCost int64  `json:"delivery_cost" validate:"required,gte=0"`
	GoodsTotal   int64  `json:"goods_total" validate:"required,gte=0"`
	CustomFee    int64  `json:"custom_fee"`
}
