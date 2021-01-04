package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"time"
)

// ChatroomRequest represents a muc chatroom request.
// swagger:model
type ChatroomRequest struct {
	Enabled     bool	`json:"enabled"`

	// the name for the chatroom
	//
	// required: true
	Name     	string	`json:"name"`

	// the chatroom display_name
	//
	// required: true
	DisplayName	string	`json:"display_name"`

	// the chatroom type
	//
	// required: true
	Type		string	`json:"type"`

	Purpose     string	`json:"purpose,omitempty"`
	Header      string	`json:"header,omitempty"`
}

// ChatroomRequest represents a muc chatroom response.
// swagger:model
type Chatroom struct {
	// the id for the chatroom
	//
	// required: false
	// min: 1
	ID          	int		`json:"id"`
	ChatroomRequest
	CreatedAt     time.Time	`json:"created_at,omitempty"`
	UpdatedAt     time.Time	`json:"updated_at,omitempty"`
	CreatorUserID	string	`json:"created_userid,omitempty"`
}

// ErrChatroomNotFound is an error raised when a ChatroomRequest can not be found in the database
var ErrChatroomNotFound = fmt.Errorf("ChatroomRequest not found")

// Chatrooms defines a slice of ChatroomRequest
type Chatrooms []*Chatroom

type ChatroomsDB struct {
	log      hclog.Logger
}

func NewChatroomsDB(l hclog.Logger) *ChatroomsDB {
	pb := &ChatroomsDB{l}
	return pb
}

// GetChatrooms returns all Chatrooms from the database
func (p *ChatroomsDB) GetChatrooms() (Chatrooms, error) {
	pr := Chatrooms{}
	for _, p := range ChatroomList {
		np := *p
		pr = append(pr, &np)
	}

	return pr, nil
}

// GetChatroomByID returns a single ChatroomRequest which matches the id from the
// database.
// If a ChatroomRequest is not found this function returns a ChatroomNotFound error
func (p *ChatroomsDB) GetChatroomByID(id int) (*Chatroom, error) {
	i := findIndexByChatroomID(id)
	if id == -1 {
		return nil, ErrChatroomNotFound
	}

	np := *ChatroomList[i]

	return &np, nil
}

// UpdateChatroom replaces a ChatroomRequest in the database with the given
// item.
// If a ChatroomRequest with the given id does not exist in the database
// this function returns a ChatroomNotFound error
func (p *ChatroomsDB) UpdateChatroom(pr Chatroom) error {
	i := findIndexByChatroomID(pr.ID)
	if i == -1 {
		return ErrChatroomNotFound
	}

	// update the ChatroomRequest in the DB
	ChatroomList[i] = &pr

	return nil
}

// AddChatroom adds a new ChatroomRequest to the database
func (p *ChatroomsDB) AddChatroom(pr *Chatroom) {
	// get the next id in sequence
	maxID := ChatroomList[len(ChatroomList)-1].ID
	pr.ID = maxID + 1
	ChatroomList = append(ChatroomList, pr)
}

// DeleteChatroom deletes a ChatroomRequest from the database
func (p *ChatroomsDB) DeleteChatroom(id int) error {
	i := findIndexByChatroomID(id)
	if i == -1 {
		return ErrChatroomNotFound
	}

	ChatroomList = append(ChatroomList[:i], ChatroomList[i+1:]...)
	return nil
}

// findIndex finds the index of a ChatroomRequest in the database
// returns -1 when no ChatroomRequest can be found
func findIndexByChatroomID(id int) int {
	for i, p := range ChatroomList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var ChatroomList = []*Chatroom{
	&Chatroom{
		ChatroomRequest: ChatroomRequest{
			Name:"chatroom01",
			DisplayName: "Display name test1",
		},
		ID: 1,
	},
	&Chatroom{
		ChatroomRequest: ChatroomRequest{
			Name:"chatroom02",
			DisplayName: "Display name test2",
		},
		ID: 2,
	},
}

////顺序初始化
//s1 := Student{Person{"ck_god", 1, 18}, 1, "sz"}
//fmt.Printf("s1=%+v\n", s1)
//
////部分成员初始化1
//s2 := Student{Person: Person{"xiaobai", 0, 20}, id: 2, addr: "sz"}
//
////部分成员初始化2
//s3 := Student{Person: Person{name: "kavai"}, id: 3}
//fmt.Println(s2, s3)
