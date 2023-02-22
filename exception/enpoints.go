package exception

import "github.com/faridlan/product-api/model/web"

func EndpointsAdmin() []web.Enpoint {

	return []web.Enpoint{
		// {
		// 	Url:    "/api/products/",
		// 	Method: "POST",
		// },
		{
			Url:    "/api/seeder/products/",
			Method: "POST",
		},
		// {
		// 	Url:    "/api/products/:id",
		// 	Method: "PUT",
		// },
		// {
		// 	Url:    "/api/products/:id",
		// 	Method: "DELETE",
		// },
		{
			Url:    "/api/seeder/products/",
			Method: "DELETE",
		},
		// {
		// 	Url:    "/api/products/:id",
		// 	Method: "GET",
		// },
		// {
		// 	Url:    "/api/products/",
		// 	Method: "GET",
		// },
		{
			Url:    "/api/log/",
			Method: "GET",
		},
	}

}
func EndpointsGlobal() []web.Enpoint {

	return []web.Enpoint{
		{
			Url:    "/api/products/",
			Method: "POST",
		},
		{
			Url:    "/api/products/:id",
			Method: "PUT",
		},
		{
			Url:    "/api/products/:id",
			Method: "DELETE",
		},
		{
			Url:    "/api/products/:id",
			Method: "GET",
		},
		{
			Url:    "/api/products/",
			Method: "GET",
		},
	}

}
