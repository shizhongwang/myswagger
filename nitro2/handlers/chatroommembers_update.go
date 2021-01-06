package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"net/http"
	"time"
)

// swagger:route PUT /chatroomMembers chatroomMembers updateChatroomMember
// Update a chatroomMembers details
//
// responses:
//	200: chatroomMemberResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update chatroomMembers
func (p *ChatroomMembers) Update(c *gin.Context) {
	var json data.ChatroomMember
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UpdatedAt = time.Now()
	err := p.ChatroomMemberDB.UpdateChatroomMember(json)
	if err == data.ErrChatroomMemberNotFound {
		p.l.Error("Unable to fetch chatroom member", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)
}

// swagger:route PATCH /chatroommembers/{chatroomid}/:userid chatroomMembers updateChatroomMemberLastViewAt
// Update a chatroomMembers LastViewedAt field
//
// responses:
//	200: chatroomMemberResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update chatroomMembers
func (p *ChatroomMembers) UpdateLastViewedAt(c *gin.Context) {
	chatroomid := c.Param("chatroomid")
	userid := c.Param("userid")

	var json data.ChatroomMemberLastViewedAt
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, err := p.ChatroomMemberDB.UpdateLastViewedAt(chatroomid, userid, json.LastViewedAt)
	if err == data.ErrChatroomMemberNotFound {
		p.l.Error("Unable to fetch chatroom member", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, rep)
}
