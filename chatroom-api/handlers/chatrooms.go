package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/shizhongwang/myswagger/Chatroom-api/data"
)

// KeyChatroom is a key used for the Chatroom object in the context
type KeyChatroom struct{}

// Chatrooms handler for getting and updating Chatrooms
type Chatrooms struct {
	l         hclog.Logger
	ChatroomDB *data.ChatroomsDB
}

// NewChatrooms returns a new Chatrooms handler with the given logger
func NewChatrooms(l hclog.Logger, pdb *data.ChatroomsDB) *Chatrooms {
	return &Chatrooms{l, pdb}
}

// ErrInvalidChatroomPath is an error message when the Chatroom path is not valid
var ErrInvalidChatroomPath = fmt.Errorf("Invalid Path, path should be /Chatrooms/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getChatroomID returns the Chatroom ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getChatroomID(r *http.Request) int {
	// parse the Chatroom id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
