package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"time"
)

// ChatroomMemeberRequest represents a muc chatroom request.
// swagger:model
type ChatroomMemeberRequest struct {
	// the id for the user
	//
	// required: true
	UserID     	string	`json:"user_id"`

	// the user role in the chatroom
	//
	// required: failse
	Role	string	`json:"role"`
}

// ChatroomMemeberRequest represents a muc chatroom response.
// swagger:model
type ChatroomMemeber struct {
	// the id for the chatroom
	//
	// required: false
	// min: 1
	ChatroomID          	string		`json:"chatroom_id"`
	ChatroomMemeberRequest
	UpdatedAt     time.Time	`json:"updated_at,omitempty"`
}

// ErrChatroomMemeberNotFound is an error raised when a ChatroomMemeberRequest can not be found in the database
var ErrChatroomMemeberNotFound = fmt.Errorf("ChatroomMemeberRequest not found")

// ChatroomMemebers defines a slice of ChatroomMemeberRequest
type ChatroomMemebers []*ChatroomMemeber

type ChatroomMemebersDB struct {
	log      hclog.Logger
}

func NewChatroomMemebersDB(l hclog.Logger) *ChatroomMemebersDB {
	pb := &ChatroomMemebersDB{l}
	return pb
}

// GetChatroomMemebers returns all ChatroomMemebers from the database
func (p *ChatroomMemebersDB) GetChatroomMemebers() (ChatroomMemebers, error) {
	pr := ChatroomMemebers{}
	for _, p := range ChatroomMemeberList {
		np := *p
		pr = append(pr, &np)
	}

	return pr, nil
}

// GetChatroomMemeberByID returns a single ChatroomMemeberRequest which matches the id from the
// database.
// If a ChatroomMemeberRequest is not found this function returns a ChatroomMemeberNotFound error
func (p *ChatroomMemebersDB) GetChatroomMemebersByChatroomID(chatroomID string) (ChatroomMemebers, error) {
	pr := ChatroomMemebers{}
	for _, p := range ChatroomMemeberList {
		np := *p
		if np.ChatroomID == chatroomID {
			pr = append(pr, &np)
		}
	}

	return pr, nil
}

// UpdateChatroomMemeber replaces a ChatroomMemeberRequest in the database with the given
// item.
// If a ChatroomMemeberRequest with the given id does not exist in the database
// this function returns a ChatroomMemeberNotFound error
func (p *ChatroomMemebersDB) UpdateChatroomMemeber(pr ChatroomMemeber) error {
	i := 1
	// update the ChatroomMemeberRequest in the DB
	ChatroomMemeberList[i] = &pr

	return nil
}

// AddChatroomMemeber adds a new ChatroomMemeberRequest to the database
func (p *ChatroomMemebersDB) AddChatroomMemeber(pr *ChatroomMemeber) {
	ChatroomMemeberList = append(ChatroomMemeberList, pr)
}

// DeleteChatroomMemeber deletes a ChatroomMemeberRequest from the database
func (p *ChatroomMemebersDB) DeleteChatroomMemeber(chatroonID string, userID string) error {
	i := 1
	ChatroomMemeberList = append(ChatroomMemeberList[:i], ChatroomMemeberList[i+1:]...)
	return nil
}

var ChatroomMemeberList = []*ChatroomMemeber{
	&ChatroomMemeber{
		ChatroomMemeberRequest: ChatroomMemeberRequest{
			UserID:"1",
			Role: "speaker",
		},
		ChatroomID: "1",
	},
	&ChatroomMemeber{
		ChatroomMemeberRequest: ChatroomMemeberRequest{
			UserID:"2",
			Role: "listener",
		},
		ChatroomID: "2",
	},
}
