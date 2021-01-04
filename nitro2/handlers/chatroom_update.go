package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"net/http"
	"time"
)

// swagger:route PUT /chatrooms chatrooms updateChatroom
// Update a chatrooms details
//
// responses:
//	200: chatroomResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update chatrooms
func (p *Chatrooms) Update(c *gin.Context) {
	var json data.Chatroom
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UpdatedAt = time.Now()
	json.CreatorUserID = "test userid"

	err := p.ChatroomDB.UpdateChatroom(json)
	if err == data.ErrChatroomNotFound {
		p.l.Error("Unable to fetch chatroom", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}
