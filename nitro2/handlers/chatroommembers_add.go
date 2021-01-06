package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shizhongwang/myswagger/nitro2/data"
	"io/ioutil"
	"net/http"
	"time"
)

// swagger:route POST /chatroomMembers{chatroomid} chatroomMembers createChatroomMembers
// AddMember a new chatroomMember
//
// responses:
//	200: chatroomMemberResponse
//  422: errorValidation
//  501: errorResponse
// AddMember handles POST requests to add new chatroomMembers
func (p *ChatroomMembers) AddMember(c *gin.Context) {
	//body中的内容 ioutil.ReadAll 读取过就不存在了, need to re-write into body
	b, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	p.l.Info("create chatroomMember" + string(b[:]))

	var json data.ChatroomMemberRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.l.Debug("Inserting chatroomMember: %#v\n", json)

	chatroomid := c.Param("chatroomid")
	cr := &data.ChatroomMember{
		ChatroomMemberRequest: json,
		ChatroomID: chatroomid,
		UpdatedAt:  time.Now(),
	}

	p.ChatroomMemberDB.AddChatroomMember(cr)
	c.JSON(http.StatusOK, cr)
}


func (p *ChatroomMembers) AddOrGetMemberByID(c *gin.Context) {
	if c.Request.Method == "POST" {
		p.AddMember(c)
	} else if c.Request.Method == "GET" {
		p.GetMembersByChatroomID(c)
	}
}
