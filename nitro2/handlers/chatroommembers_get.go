package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"net/http"
)

// swagger:route GET /chatroomMembers chatroomMembers listChatroomMembersAll
// Return a list of ChatroomMembers from the database
// responses:
//	200: chatroomMembersResponse

// GetChatroomsAll handles GET requests and returns all current ChatroomMembers
func (p *ChatroomMembers) ListAll(c *gin.Context) {
	prods, err := p.ChatroomMemberDB.GetChatroomMembers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prods)
}

// swagger:route GET /chatroomMembers/{chatroomid} chatroomMembers listChatroomMembersByChatroom
// Return a list of ChatroomMembers from the database
// responses:
//	200: chatroomMembersResponse
//	404: errorResponse

// GetMembersByChatroomID handles GET requests
func (p *ChatroomMembers) GetMembersByChatroomID(c *gin.Context) {
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


// swagger:route GET /chatroomMembers/{userid} chatroomMembers listChatroomMembersByUserID
// Return a list of ChatroomMembers from the database
// responses:
//	200: chatroomMembersResponse
//	404: errorResponse

// GetMembersByChatroomID handles GET requests
func (p *ChatroomMembers) GetMembersByUserID(c *gin.Context) {
	userid := c.Param("userid")
	p.l.Debug("Get record", "userid: ", userid)
	prod, err := p.ChatroomMemberDB.GetChatroomMembersByUserID(userid)

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
