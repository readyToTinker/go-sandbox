package handlers

import (
	"net/http"
)

func New() http.Handler {

	http.HandleFunc("GET /hello", helloHandler)

	http.HandleFunc("GET /api/products", nil)     // list all the available products
	http.HandleFunc("GET /api/product/{id}", nil) // get the details of one product
	http.HandleFunc("POST /api/product", nil)     // add new product
	http.HandleFunc("POST /api/order", nil)       // start an order which contains the amount and product id, returns an order id
	http.HandleFunc("GET /api/order/{id}", nil)   // check the status of the order based on order id

	return nil
}

func helloHandler(wr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	wr.Write([]byte("Hello World"))
}
