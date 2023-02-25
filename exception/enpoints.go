package exception

import "github.com/faridlan/product-api/model/web"

func EndpointsGlobal() []web.Enpoint {
	return []web.Enpoint{
		{
			Url:    "/api/login",
			Method: "POST",
		},
		{
			Url:    "/api/products",
			Method: "GET",
		},
		{
			Url:    "/api/products/:id",
			Method: "GET",
		},
	}

}
