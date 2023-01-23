package domain

type Product struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Price     int    `json:"price,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
