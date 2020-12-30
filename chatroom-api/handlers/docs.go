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
	"github.com/shizhongwang/myswagger/chatroom-api/data"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

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

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
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
