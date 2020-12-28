package data

import "time"

// Chatroom represents a muc chatroom.
// swagger:model
type Chatroom struct {
	// the id for the chatroom
	//
	// required: false
	// min: 1
	ID          int

	URI         string
	Enabled     bool
	Deleted     bool
	Public      bool
	Visible     bool
	Topic       string
	EnableUBM   bool
	Created     time.Time
	Updated     time.Time
	DisplayName string
}

//"team_id": "string",
//"name": "string",
//"display_name": "string",
//"purpose": "string",
//"header": "string",
//"type": "string"