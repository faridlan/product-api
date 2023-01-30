package web

type ProductCreate struct {
	Name      string `json:"name,omitempty" validate:"required"`
	Price     int    `json:"price,omitempty" validate:"required,numeric"`
	Quantity  int    `json:"quantity,omitempty" validate:"required,numeric"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
