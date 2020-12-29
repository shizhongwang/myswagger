package handlers

import (
	"net/http"

	"github.com/shizhongwang/myswagger/chatroom-api/data"
)

// swagger:route POST /chatrooms chatrooms createChatroom
// Create a new chatroom
//
// responses:
//	200: chatroomResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new chatrooms
func (p *Chatrooms) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the chatroom from the context
	prod := r.Context().Value(KeyChatroom{}).(data.ChatroomRequest)

	p.l.Debug("Inserting chatroom: %#v\n", prod)
	p.ChatroomDB.AddChatroom(prod)

	p.ChatroomDB.AddChatroom( Chatroom: Chatroom{
		ChatroomRequest: ChatroomRequest{Name:"chatroom02"},
	},)
}
