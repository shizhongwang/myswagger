package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/chatroom-api/data"
	"net/http"
	"strconv"
)

// swagger:route GET /chatrooms chatrooms listChatrooms
// Return a list of Chatrooms from the database
// responses:
//	200: chatroomsResponse

// ListAll handles GET requests and returns all current Chatrooms
func (p *Chatrooms) ListAll(c *gin.Context) {
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

// ListSingle handles GET requests
func (p *Chatrooms) ListSingle(c *gin.Context) {
	id := c.Param("id")
	p.l.Debug("Get record", "id", id)
	idint, err := strconv.Atoi(id)
	prod, err := p.ChatroomDB.GetChatroomByID(idint)

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
