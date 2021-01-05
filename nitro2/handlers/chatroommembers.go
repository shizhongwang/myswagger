package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/shizhongwang/myswagger/nitro2/data"
)

// ChatroomMembers handler for getting and updating ChatroomMembers
type ChatroomMembers struct {
	l         hclog.Logger
	ChatroomMemberDB *data.ChatroomMembersDB
}

// NewChatroomMembers returns a new ChatroomMembers handler with the given logger
func NewChatroomMembers(l hclog.Logger, pdb *data.ChatroomMembersDB) *ChatroomMembers {
	return &ChatroomMembers{l, pdb}
}