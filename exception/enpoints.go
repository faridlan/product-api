package exception

import "github.com/faridlan/product-api/model/web"

func EndpointsGlobal() []web.Enpoint {
	return []web.Enpoint{
		{
			Url:    "/api/seeder/products",
			Method: "POST",
		},
		{
			Url:    "/api/seeder/products",
			Method: "DELETE",
		},
		{
			Url:    "/api/log",
			Method: "GET",
		},
	}

}
