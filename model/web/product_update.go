package web

import "github.com/faridlan/product-api/helper/mysql"

type ProductUpdate struct {
	Id        int           `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Price     int           `json:"price,omitempty"`
	Quantity  int           `json:"quantity,omitempty"`
	UpdatedAt mysql.NullInt `json:"updated_at,omitempty"`
}
