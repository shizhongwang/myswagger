package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/chatroom-api/data"
	"net/http"
	"strconv"
)

// swagger:route DELETE /chatrooms/{id} chatrooms deleteChatroom
// Update a chatrooms details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Chatrooms) Delete(c *gin.Context) {
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)
	p.l.Debug("Deleting record", "id", id)

	err := p.ChatroomDB.DeleteChatroom(idint)
	if err == data.ErrChatroomNotFound {
		p.l.Error("Unable to delete record id does not exist", "error", err)
		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	}

	if err != nil {
		p.l.Error("Unable to delete record", "error", err)

		c.JSON(http.StatusInternalServerError, &GenericError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "OK")
}

