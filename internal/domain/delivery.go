package domain

type DeliveryInfo struct {
	Name    string `json:"name" validate:"required,gte=4"`
	Phone   string `json:"phone" validate:"required,gte=10"`
	Zip     string `json:"zip" validate:"required,gte=5"`
	City    string `json:"city" validate:"required,gte=6"`
	Address string `json:"address" validate:"required,gte=6"`
	Region  string `json:"region" validate:"required,gte=4"`
	Email   string `json:"email" validate:"required,email"`
}
