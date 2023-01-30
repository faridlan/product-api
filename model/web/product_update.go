package web

type ProductUpdate struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty" validate:"required"`
	Price     int    `json:"price,omitempty" validate:"required"`
	Quantity  int    `json:"quantity,omitempty" validate:"required"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
