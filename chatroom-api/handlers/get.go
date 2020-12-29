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
	cur := c.Request.URL.Query().Get("currency")

	prods, err := p.ChatroomDB.GetChatrooms(cur)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prods)
}

//func (p *Chatrooms) ListAll(rw http.ResponseWriter, r *http.Request) {
//	p.l.Debug("Get all records")
//	rw.Header().Add("Content-Type", "application/json")
//
//	cur := r.URL.Query().Get("currency")
//
//	prods, err := p.ChatroomDB.GetChatrooms(cur)
//	if err != nil {
//		rw.WriteHeader(http.StatusInternalServerError)
//		data.ToJSON(&GenericError{Message: err.Error()}, rw)
//		return
//	}
//
//	err = data.ToJSON(prods, rw)
//	if err != nil {
//		// we should never be here but log the error just incase
//		p.l.Error("Unable to serializing chatroom", "error", err)
//	}
//}

// swagger:route GET /chatrooms/{id} chatrooms listSingleChatroom
// Return a list of Chatrooms from the database
// responses:
//	200: chatroomResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Chatrooms) ListSingle(c *gin.Context) {
	//id := getChatroomID(c.Request)
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
//
//// ListSingle handles GET requests
//func (p *Chatrooms) ListSingle(rw http.ResponseWriter, r *http.Request) {
//	rw.Header().Add("Content-Type", "application/json")
//
//	id := getChatroomID(r)
//	//cur := r.URL.Query().Get("currency")
//
//	p.l.Debug("Get record", "id", id)
//
//	prod, err := p.ChatroomDB.GetChatroomByID(id)
//
//	switch err {
//	case nil:
//
//	case data.ErrChatroomNotFound:
//		p.l.Error("Unable to fetch chatroom", "error", err)
//
//		rw.WriteHeader(http.StatusNotFound)
//		data.ToJSON(&GenericError{Message: err.Error()}, rw)
//		return
//	default:
//		p.l.Error("Unable to fetching chatroom", "error", err)
//
//		rw.WriteHeader(http.StatusInternalServerError)
//		data.ToJSON(&GenericError{Message: err.Error()}, rw)
//		return
//	}
//
//	err = data.ToJSON(prod, rw)
//	if err != nil {
//		// we should never be here but log the error just incase
//		p.l.Error("Unable to serializing chatroom", err)
//	}
//}
