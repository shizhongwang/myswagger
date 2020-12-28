package data

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	protos "github.com/shizhongwang/myswagger/currency/protos/currency"
)

// Chatroom represents a muc chatroom.
// swagger:model
type Chatroom struct {
	// the id for the chatroom
	//
	// required: false
	// min: 1
	ID          int
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


// ErrChatroomNotFound is an error raised when a Chatroom can not be found in the database
var ErrChatroomNotFound = fmt.Errorf("Chatroom not found")

// Chatrooms defines a slice of Chatroom
type Chatrooms []*Chatroom

type ChatroomsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
	rates    map[string]float64
	client   protos.Currency_SubscribeRatesClient
}

func NewChatroomsDB(c protos.CurrencyClient, l hclog.Logger) *ChatroomsDB {
	pb := &ChatroomsDB{c, l, make(map[string]float64), nil}

	go pb.handleUpdates()

	return pb
}

func (p *ChatroomsDB) handleUpdates() {
	sub, err := p.currency.SubscribeRates(context.Background())
	if err != nil {
		p.log.Error("Unable to subscribe for rates", "error", err)
	}

	p.client = sub

	for {
		rr, err := sub.Recv()
		p.log.Info("Recieved updated rate from server", "dest", rr.GetDestination().String())

		if err != nil {
			p.log.Error("Error receiving message", "error", err)
			return
		}

		p.rates[rr.Destination.String()] = rr.Rate
	}
}

// GetChatrooms returns all Chatrooms from the database
func (p *ChatroomsDB) GetChatrooms(currency string) (Chatrooms, error) {
	if currency == "" {
		return ChatroomList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	pr := Chatrooms{}
	for _, p := range ChatroomList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return pr, nil
}

// GetChatroomByID returns a single Chatroom which matches the id from the
// database.
// If a Chatroom is not found this function returns a ChatroomNotFound error
func (p *ChatroomsDB) GetChatroomByID(id int, currency string) (*Chatroom, error) {
	i := findIndexByChatroomID(id)
	if id == -1 {
		return nil, ErrChatroomNotFound
	}

	if currency == "" {
		return ChatroomList[i], nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	np := *ChatroomList[i]
	np.Price = np.Price * rate

	return &np, nil
}

// UpdateChatroom replaces a Chatroom in the database with the given
// item.
// If a Chatroom with the given id does not exist in the database
// this function returns a ChatroomNotFound error
func (p *ChatroomsDB) UpdateChatroom(pr Chatroom) error {
	i := findIndexByChatroomID(pr.ID)
	if i == -1 {
		return ErrChatroomNotFound
	}

	// update the Chatroom in the DB
	ChatroomList[i] = &pr

	return nil
}

// AddChatroom adds a new Chatroom to the database
func (p *ChatroomsDB) AddChatroom(pr Chatroom) {
	// get the next id in sequence
	maxID := ChatroomList[len(ChatroomList)-1].ID
	pr.ID = maxID + 1
	ChatroomList = append(ChatroomList, &pr)
}

// DeleteChatroom deletes a Chatroom from the database
func (p *ChatroomsDB) DeleteChatroom(id int) error {
	i := findIndexByChatroomID(id)
	if i == -1 {
		return ErrChatroomNotFound
	}

	ChatroomList = append(ChatroomList[:i], ChatroomList[i+1])

	return nil
}

// findIndex finds the index of a Chatroom in the database
// returns -1 when no Chatroom can be found
func findIndexByChatroomID(id int) int {
	for i, p := range ChatroomList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func (p *ChatroomsDB) getRate(destination string) (float64, error) {
	// if cached return
	if r, ok := p.rates[destination]; ok {
		return r, nil
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	// get initial rate
	resp, err := p.currency.GetRate(context.Background(), rr)
	p.rates[destination] = resp.Rate // update cache

	// subscribe for updates
	p.client.Send(rr)

	return resp.Rate, err
}

var ChatroomList = []*Chatroom{
	&Chatroom{
		ID:          1,
	},
	&Chatroom{
		ID:          2,
	},
}
