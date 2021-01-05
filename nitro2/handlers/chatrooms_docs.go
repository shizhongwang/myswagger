// Package classification of ChatroomRequest API
//
// Documentation for ChatroomRequest API
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

// A list of chatrooms
// swagger:response chatroomsResponse
type chatroomsResponseWrapper struct {
	// All current chatrooms
	// in: body
	Body []data.ChatroomRequest
}

// Data structure representing a single chatroom
// swagger:response chatroomResponse
type chatroomResponseWrapper struct {
	// Newly created chatroom
	// in: body
	Body data.Chatroom
}


// swagger:parameters createChatroom
type chatroomRequestParamsWrapper struct {
	// ChatroomRequest data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.ChatroomRequest
}

// Data structure representing a single chatroom
// swagger:parameters updateChatroom
type chatroomParamsWrapper struct {
	// Newly created chatroom
	// in: body
	Body data.Chatroom
}

// swagger:parameters listChatrooms listSingleChatroom
type chatroomQueryParam struct {
	// Currency used when returning the price of the chatroom,
	// when not specified currency is returned in GBP.
	// in: query
	// required: false
	Currency string
}

// swagger:parameters listSingleChatroom deleteChatroom
type chatroomIDParamsWrapper struct {
	// The id of the chatroom for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
