package web

type ProductUpdate struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Price     int    `json:"price,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
