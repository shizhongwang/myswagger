package handlers

import (
	"net/http"

	"github.com/shizhongwang/myswagger/chatroom-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Chatrooms) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyChatroom{}).(data.Chatroom)

	p.l.Debug("Inserting product: %#v\n", prod)
	p.ChatroomDB.AddChatroom(prod)
}
