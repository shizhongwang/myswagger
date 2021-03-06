package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"time"
)

// ChatroomMemberRequest represents a muc chatroom request.
// swagger:model
type ChatroomMemberRequest struct {
	// the id for the user
	//
	// required: true
	UserID     	string	`json:"user_id"`

	// the user role in the chatroom
	//
	// required: failse
	Role	string	`json:"role"`
}

// ChatroomMemberLastViewedAt represents a LastViewedAt update body
// swagger:model
type ChatroomMemberLastViewedAt struct {
	LastViewedAt time.Time	`json:"lastviewed_at,omitempty"`
}

// ChatroomMemberRequest represents a muc chatroom response.
// swagger:model
type ChatroomMember struct {
	// the id for the chatroom
	//
	// required: false
	// min: 1
	ChatroomID          	string		`json:"chatroom_id"`
	ChatroomMemberRequest
	UpdatedAt     time.Time	`json:"updated_at,omitempty"`
	LastViewedAt	time.Time	`json:"lastviewed_at,omitempty"`
}

// ErrChatroomMemberNotFound is an error raised when a ChatroomMemberRequest can not be found in the database
var ErrChatroomMemberNotFound = fmt.Errorf("ChatroomMemberRequest not found")

// ChatroomMembers defines a slice of ChatroomMemberRequest
type ChatroomMembers []*ChatroomMember

type ChatroomMembersDB struct {
	log      hclog.Logger
}

func NewChatroomMembersDB(l hclog.Logger) *ChatroomMembersDB {
	pb := &ChatroomMembersDB{l}
	return pb
}

// GetChatroomMembers returns all ChatroomMembers from the database
func (p *ChatroomMembersDB) GetChatroomMembers() (ChatroomMembers, error) {
	pr := ChatroomMembers{}
	for _, p := range ChatroomMemberList {
		np := *p
		pr = append(pr, &np)
	}

	return pr, nil
}

// GetChatroomMembersByChatroomID returns a single ChatroomMemberRequest which matches the id from the
// database.
// If a ChatroomMemberRequest is not found this function returns a ChatroomMemberNotFound error
func (p *ChatroomMembersDB) GetChatroomMembersByChatroomID(chatroomID string) (ChatroomMembers, error) {
	pr := ChatroomMembers{}
	for _, p := range ChatroomMemberList {
		np := *p
		if np.ChatroomID == chatroomID {
			pr = append(pr, &np)
		}
	}

	return pr, nil
}

// GetChatroomMembersByUserID returns a single ChatroomMemberRequest which matches the id from the
// database.
// If a ChatroomMemberRequest is not found this function returns a ChatroomMemberNotFound error
func (p *ChatroomMembersDB) GetChatroomMembersByUserID(userID string) (ChatroomMembers, error) {
	pr := ChatroomMembers{}
	for _, p := range ChatroomMemberList {
		np := *p
		if np.UserID == userID {
			pr = append(pr, &np)
		}
	}

	return pr, nil
}


// UpdateChatroomMember replaces a ChatroomMemberRequest in the database with the given
// item.
// If a ChatroomMemberRequest with the given id does not exist in the database
// this function returns a ChatroomMemberNotFound error
func (p *ChatroomMembersDB) UpdateLastViewedAt(chatroomid string, userid string, lastviewedat time.Time) (*ChatroomMember, error) {
	i := 1
	// update the ChatroomMemberRequest in the DB
	ChatroomMemberList[i].LastViewedAt = lastviewedat

	return ChatroomMemberList[i], nil
}

// UpdateChatroomMember replaces a ChatroomMemberRequest in the database with the given
// item.
// If a ChatroomMemberRequest with the given id does not exist in the database
// this function returns a ChatroomMemberNotFound error
func (p *ChatroomMembersDB) UpdateChatroomMember(pr ChatroomMember) error {
	i := 1
	// update the ChatroomMemberRequest in the DB
	ChatroomMemberList[i] = &pr

	return nil
}

// AddChatroomMember adds a new ChatroomMemberRequest to the database
func (p *ChatroomMembersDB) AddChatroomMember(pr *ChatroomMember) {
	ChatroomMemberList = append(ChatroomMemberList, pr)
}

// DeleteChatroomMember deletes a ChatroomMemberRequest from the database
func (p *ChatroomMembersDB) DeleteChatroomMember(chatroonID string, userID string) error {
	i := 1
	ChatroomMemberList = append(ChatroomMemberList[:i], ChatroomMemberList[i+1:]...)
	return nil
}

var ChatroomMemberList = []*ChatroomMember{
	&ChatroomMember{
		ChatroomMemberRequest: ChatroomMemberRequest{
			UserID:"1",
			Role: "speaker",
		},
		ChatroomID: "1",
		UpdatedAt: time.Now(),
		LastViewedAt: time.Now(),
	},
	&ChatroomMember{
		ChatroomMemberRequest: ChatroomMemberRequest{
			UserID:"2",
			Role: "listener",
		},
		ChatroomID: "2",
		UpdatedAt: time.Now(),
		LastViewedAt: time.Now(),
	},
}
