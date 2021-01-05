package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/shizhongwang/myswagger/nitro2/data"
)

// ChatroomMember handler for getting and updating ChatroomMember
type ChatroomMember struct {
	l         hclog.Logger
	ChatroomMemberDB *data.ChatroomMemberDB
}

// NewChatroomMembers returns a new ChatroomMember handler with the given logger
func NewChatroomMembers(l hclog.Logger, pdb *data.ChatroomMemberDB) *ChatroomMember {
	return &ChatroomMember{l, pdb}
}