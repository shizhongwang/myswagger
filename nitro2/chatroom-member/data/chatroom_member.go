package data
//
//import (
//	"fmt"
//	"github.com/hashicorp/go-hclog"
//	"time"
//)
//
////`chatroom_id` varchar(26) NOT NULL,
////`user_id` varchar(26) NOT NULL,
////`online` tinyint(1) NOT NULL,
////`role` enum('moderator','none','participant','visitor') NOT NULL DEFAULT 'none',
////`lastviewed_at` timestamp DEFAULT NULL,
////`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//
//
//// ChatroomMemeberMemeberRequest represents a muc chatroom request.
//// swagger:model
//type ChatroomMemeberMemeberRequest struct {
//	Enabled     bool	`json:"enabled"`
//
//	// the name for the chatroom
//	//
//	// required: true
//	Name     	string	`json:"name"`
//
//	// the chatroom display_name
//	//
//	// required: true
//	DisplayName	string	`json:"display_name"`
//
//	// the chatroom type
//	//
//	// required: true
//	Type		string	`json:"type"`
//
//	Purpose     string	`json:"purpose,omitempty"`
//	Header      string	`json:"header,omitempty"`
//}
//
//// ChatroomMemeberMemeberRequest represents a muc chatroom response.
//// swagger:model
//type ChatroomMemeberMemeber struct {
//	// the id for the chatroom
//	//
//	// required: false
//	// min: 1
//	ID          	int		`json:"id"`
//	ChatroomMemeberMemeberRequest
//	CreatedAt     time.Time	`json:"created_at,omitempty"`
//	UpdatedAt     time.Time	`json:"updated_at,omitempty"`
//	CreatorUserID	string	`json:"created_userid,omitempty"`
//}
//
//// ErrChatroomMemeberMemeberNotFound is an error raised when a ChatroomMemeberMemeberRequest can not be found in the database
//var ErrChatroomMemeberMemeberNotFound = fmt.Errorf("ChatroomMemeberMemeberRequest not found")
//
//// ChatroomMemeberMemebers defines a slice of ChatroomMemeberMemeberRequest
//type ChatroomMemeberMemebers []*ChatroomMemeberMemeber
//
//type ChatroomMemeberMemebersDB struct {
//	log      hclog.Logger
//}
//
//func NewChatroomMemeberMemebersDB(l hclog.Logger) *ChatroomMemeberMemebersDB {
//	pb := &ChatroomMemeberMemebersDB{l}
//	return pb
//}
//
//// GetChatroomMemeberMemebers returns all ChatroomMemeberMemebers from the database
//func (p *ChatroomMemeberMemebersDB) GetChatroomMemeberMemebers() (ChatroomMemeberMemebers, error) {
//	pr := ChatroomMemeberMemebers{}
//	for _, p := range ChatroomMemeberMemeberList {
//		np := *p
//		pr = append(pr, &np)
//	}
//
//	return pr, nil
//}
//
//// GetChatroomMemeberMemeberByID returns a single ChatroomMemeberMemeberRequest which matches the id from the
//// database.
//// If a ChatroomMemeberMemeberRequest is not found this function returns a ChatroomMemeberMemeberNotFound error
//func (p *ChatroomMemeberMemebersDB) GetChatroomMemeberMemeberByID(id int) (*ChatroomMemeberMemeber, error) {
//	i := findIndexByChatroomMemeberMemeberID(id)
//	if id == -1 {
//		return nil, ErrChatroomMemeberMemeberNotFound
//	}
//
//	np := *ChatroomMemeberMemeberList[i]
//
//	return &np, nil
//}
//
//// UpdateChatroomMemeberMemeber replaces a ChatroomMemeberMemeberRequest in the database with the given
//// item.
//// If a ChatroomMemeberMemeberRequest with the given id does not exist in the database
//// this function returns a ChatroomMemeberMemeberNotFound error
//func (p *ChatroomMemeberMemebersDB) UpdateChatroomMemeberMemeber(pr ChatroomMemeberMemeber) error {
//	i := findIndexByChatroomMemeberMemeberID(pr.ID)
//	if i == -1 {
//		return ErrChatroomMemeberMemeberNotFound
//	}
//
//	// update the ChatroomMemeberMemeberRequest in the DB
//	ChatroomMemeberMemeberList[i] = &pr
//
//	return nil
//}
//
//// AddChatroomMemeberMemeber adds a new ChatroomMemeberMemeberRequest to the database
//func (p *ChatroomMemeberMemebersDB) AddChatroomMemeberMemeber(pr *ChatroomMemeberMemeber) {
//	// get the next id in sequence
//	maxID := ChatroomMemeberMemeberList[len(ChatroomMemeberMemeberList)-1].ID
//	pr.ID = maxID + 1
//	ChatroomMemeberMemeberList = append(ChatroomMemeberMemeberList, pr)
//}
//
//// DeleteChatroomMemeberMemeber deletes a ChatroomMemeberMemeberRequest from the database
//func (p *ChatroomMemeberMemebersDB) DeleteChatroomMemeberMemeber(id int) error {
//	i := findIndexByChatroomMemeberMemeberID(id)
//	if i == -1 {
//		return ErrChatroomMemeberMemeberNotFound
//	}
//
//	ChatroomMemeberMemeberList = append(ChatroomMemeberMemeberList[:i], ChatroomMemeberMemeberList[i+1:]...)
//	return nil
//}
//
//// findIndex finds the index of a ChatroomMemeberMemeberRequest in the database
//// returns -1 when no ChatroomMemeberMemeberRequest can be found
//func findIndexByChatroomMemeberMemeberID(id int) int {
//	for i, p := range ChatroomMemeberMemeberList {
//		if p.ID == id {
//			return i
//		}
//	}
//
//	return -1
//}
//
//var ChatroomMemeberMemeberList = []*ChatroomMemeberMemeber{
//	&ChatroomMemeberMemeber{
//		ChatroomMemeberMemeberRequest: ChatroomMemeberMemeberRequest{
//			Name:"chatroom01",
//			DisplayName: "Display name test1",
//		},
//		ID: 1,
//	},
//	&ChatroomMemeberMemeber{
//		ChatroomMemeberMemeberRequest: ChatroomMemeberMemeberRequest{
//			Name:"chatroom02",
//			DisplayName: "Display name test2",
//		},
//		ID: 2,
//	},
//}
//
//////顺序初始化
////s1 := Student{Person{"ck_god", 1, 18}, 1, "sz"}
////fmt.Printf("s1=%+v\n", s1)
////
//////部分成员初始化1
////s2 := Student{Person: Person{"xiaobai", 0, 20}, id: 2, addr: "sz"}
////
//////部分成员初始化2
////s3 := Student{Person: Person{name: "kavai"}, id: 3}
////fmt.Println(s2, s3)
