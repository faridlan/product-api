package web

type ProductResponse struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Price     int    `json:"price,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
