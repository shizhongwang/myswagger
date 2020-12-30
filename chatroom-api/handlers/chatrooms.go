package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/shizhongwang/myswagger/chatroom-api/data"
)

// Chatrooms handler for getting and updating Chatrooms
type Chatrooms struct {
	l         hclog.Logger
	ChatroomDB *data.ChatroomsDB
}

// NewChatrooms returns a new Chatrooms handler with the given logger
func NewChatrooms(l hclog.Logger, pdb *data.ChatroomsDB) *Chatrooms {
	return &Chatrooms{l, pdb}
}