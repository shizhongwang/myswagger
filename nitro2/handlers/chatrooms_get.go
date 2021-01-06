package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"net/http"
)

// swagger:route GET /chatrooms chatrooms listChatrooms
// Return a list of Chatrooms from the database
// responses:
//	200: chatroomsResponse

// GetChatroomsAll handles GET requests and returns all current Chatrooms
func (p *Chatrooms) GetChatroomsAll(c *gin.Context) {
	prods, err := p.ChatroomDB.GetChatrooms()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prods)
}

// swagger:route GET /chatrooms/{id} chatrooms listSingleChatroom
// Return a list of Chatrooms from the database
// responses:
//	200: chatroomResponse
//	404: errorResponse

// GetChatroomByID handles GET requests
func (p *Chatrooms) GetChatroomByID(c *gin.Context) {
	id := c.Param("id")
	p.l.Debug("Get record", "id", id)
	prod, err := p.ChatroomDB.GetChatroomByID(id)

	switch err {
	case nil:
	case data.ErrChatroomNotFound:
		p.l.Error("Unable to fetch chatroom", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	default:
		p.l.Error("Unable to fetching chatroom", "error", err)
		c.JSON(http.StatusInternalServerError, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, prod)
}
