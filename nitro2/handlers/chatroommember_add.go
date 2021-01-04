package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"io/ioutil"
	"net/http"
	"time"
)


// swagger:route POST /chatrooms chatrooms createChatroom
// Create a new chatroom
//
// responses:
//	200: chatroomResponse
//  422: errorValidation
//  501: errorResponse
// Create handles POST requests to add new chatrooms
func (p *ChatroomMember) Create(c *gin.Context) {
	//body中的内容 ioutil.ReadAll 读取过就不存在了, need to re-write into body
	b, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	p.l.Info("create chatroom" + string(b[:]))

	var json data.ChatroomRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.l.Debug("Inserting chatroom: %#v\n", json)

	cr := &data.Chatroom{
		ChatroomRequest: data.ChatroomRequest{
			Name: "chatroom02",
			Type: "normal type",
		},
		ID: "3",
		CreatedAt: time.Now(),
	}

	p.ChatroomDB.AddChatroom(cr)
	c.JSON(http.StatusOK, cr)
}
