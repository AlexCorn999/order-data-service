package domain

type ItemInfo struct {
	ChrtID      int64  `json:"chrt_id" validate:"required,gte=0"`
	TrackNumber string `json:"track_number" validate:"required,gte=10"`
	Price       int64  `json:"price" validate:"required,gte=0"`
	Rid         string `json:"rid" validate:"required,gte=5"`
	Name        string `json:"name" validate:"required,gte=4"`
	Sale        int64  `json:"sale" validate:"required,gte=0"`
	Size        string `json:"size" validate:"required,gte=0"`
	TotalPrice  int64  `json:"total_price" validate:"required,gte=0"`
	NmID        int64  `json:"nm_id" validate:"required,gte=0"`
	Brand       string `json:"brand" validate:"required,gte=4"`
	Status      int64  `json:"status" validate:"required,gte=0"`
}
