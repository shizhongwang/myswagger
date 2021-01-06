package handlers

import (
	"net/http"

	"github.com/shizhongwang/myswagger/product-api/data"
)

// swagger:route POST /products products createProduct
// AddMember a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddMember handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Debug("Inserting product: %#v\n", prod)
	p.productDB.AddProduct(prod)
}
