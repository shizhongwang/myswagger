// Package classification of ChatroomMembersRequest API
//
// Documentation for ChatroomMembersRequest API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"github.com/shizhongwang/myswagger/nitro2/data"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// A list of chatroomMembers
// swagger:response chatroomMembersResponse
type chatroomMembersResponseWrapper struct {
	// All current chatroomMembers
	// in: body
	Body data.ChatroomMembers
}

// Data structure representing a single chatroomMember
// swagger:response chatroomMemberResponse
type chatroomMemberResponseWrapper struct {
	// Newly created chatroomMember
	// in: body
	Body data.ChatroomMember
}

// swagger:parameters createChatroomMembers
type chatroomMemberRequestParamsWrapper struct {
	// ChatroomMembersRequest data structure to Update or AddMember.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.ChatroomMemberRequest
}

// Data structure representing a single chatroomMember
// swagger:parameters updateChatroomMembers
type chatroomMemberParamsWrapper struct {
	// Newly created chatroomMember
	// in: body
	Body data.ChatroomMembers
}

// swagger:parameters listChatroomMembersByUserID
type chatroomMembersByUserIDQueryParam struct {
	// By using the userid to search for the list of chatroomMembers,
	// in: path
	// required: true
	UserID string
}

// swagger:parameters listChatroomMembersByChatroom deleteChatroomMembers
type chatroomMemberIDParamsWrapper struct {
	// The id of the chatroomMember for which the operation relates
	// in: path
	// required: true
	ChatroomID string `json:"chatroomid"`
}
