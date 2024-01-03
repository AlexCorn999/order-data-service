package domain

import "time"

type Order struct {
	OrderID           string       `json:"order_uid"`
	TrackNumber       string       `json:"track_number"`
	Entry             string       `json:"entry"`
	Delivery          DeliveryInfo `json:"delivery"`
	Payment           PaymentInfo  `json:"payment"`
	Items             []Item       `json:"items"`
	Locale            string       `json:"locale"`
	InternalSignature string       `json:"internal_signature"`
	CustomerID        string       `json:"customer_id"`
	DeliveryService   string       `json:"delivery_service"`
	Shardkey          string       `json:"shardkey"`
	SmID              int64        `json:"sm_id"`
	DateCreated       time.Time    `json:"date_created"`
	OofShard          string       `json:"oof_shard"`
}
