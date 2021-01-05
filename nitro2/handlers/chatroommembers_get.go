package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"net/http"
)

// swagger:route GET /chatrooms chatrooms listChatroomMembers
// Return a list of ChatroomMembers from the database
// responses:
//	200: chatroomsResponse

// ListAll handles GET requests and returns all current ChatroomMembers
func (p *ChatroomMembers) ListAll(c *gin.Context) {
	prods, err := p.ChatroomMemberDB.GetChatroomMembers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prods)
}

// swagger:route GET /chatrooms/{id} chatrooms listSingleChatroomMember
// Return a list of ChatroomMembers from the database
// responses:
//	200: chatroomResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *ChatroomMembers) ListSingle(c *gin.Context) {
	chatroomid := c.Param("chatroomid")
	p.l.Debug("Get record", "chatroomid: ", chatroomid)
	prod, err := p.ChatroomMemberDB.GetChatroomMembersByChatroomID(chatroomid)

	switch err {
	case nil:
	case data.ErrChatroomMemberNotFound:
		p.l.Error("Unable to fetch chatroom member", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	default:
		p.l.Error("Unable to fetching chatroom member", "error", err)
		c.JSON(http.StatusInternalServerError, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, prod)
}
