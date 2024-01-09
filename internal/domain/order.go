package domain

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate           *validator.Validate
	ErrAlreadyUploaded = errors.New("the order number has already been uploaded")
	ErrIncorrectOrder  = errors.New("incorrect order id")
)

type Order struct {
	OrderID           string       `json:"order_uid" validate:"required,gte=10"`
	TrackNumber       string       `json:"track_number" validate:"required,gte=10"`
	Entry             string       `json:"entry" validate:"required,gte=4"`
	Delivery          DeliveryInfo `json:"delivery" validate:"required"`
	Payment           PaymentInfo  `json:"payment" validate:"required"`
	Items             []ItemInfo   `json:"items" validate:"required,dive,required"`
	Locale            string       `json:"locale" validate:"required,gte=2"`
	InternalSignature string       `json:"internal_signature"`
	CustomerID        string       `json:"customer_id" validate:"required,gte=2"`
	DeliveryService   string       `json:"delivery_service" validate:"required,gte=4"`
	Shardkey          string       `json:"shardkey" validate:"required"`
	SmID              int64        `json:"sm_id" validate:"required,gte=0"`
	DateCreated       time.Time    `json:"date_created" validate:"required"`
	OofShard          string       `json:"oof_shard" validate:"required"`
}

type Page struct {
	OrderID string
}

func init() {
	validate = validator.New()
}

func (o *Order) Validate() error {
	return validate.Struct(o)
}
