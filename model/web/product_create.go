package web

type ProductCreate struct {
	Name      string `json:"name,omitempty"`
	Price     int    `json:"price,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
