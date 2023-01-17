package event

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

const Yellow = "\033[33m"
const Green = "\033[32m"

const Share_Desc = `
NAME:
	share - used so that people can share the bill
	
EXAMPLE:
	share -a init birthday
	share -a add -p=10 birthday Soufiane
	share -a get birthday
	share -a split birthday
	share -a list or just share without any flags

DESCRIPTION:
	-a: stands for action & right now the action is init an event (birthday)
	-a: stands for action & right now the action is 'add' a participant in the event
	-p: stands for paid & right now the participant 'Soufiane' paid 10MAD(or $ whatever) in the 'birthday' event
	-p: if paid is not mentionned like (share -a add birthday Soufiane) then 'Soufiane' paid 0MAD
	-a: the action is 'get' participants in the event
	-a: the action is 'split' the bill of the event 'birthday'
	-a: the action is to 'list' all event with participants`

type Sharedb struct {
	db *badger.DB
}

type Event struct {
	Name         string
	Participants []Participant
}

type Participant struct {
	Name      string
	Paid      float64
	CreatedAt time.Time
}

type Transaction struct {
	Giver, Receiver Participant
	Paid            float64
}

func (e Participant) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf(Yellow+"| %-30v\t| %-40v\t| %v", e.Name, e.Paid, created+Yellow)
}

func PrintParticipants(participants []Participant) {
	for _, pt := range participants {
		fmt.Println(pt)
	}
}

func PrintPtsHeader() {
	fmt.Printf(Green+"| %-30v\t| %-40v\t| %v\n", "Name", "Amount Paid", "Created"+Green)
}

func PrintTxnsHeader() {
	fmt.Printf(Green + "*---------------Split OutPut--------------*\n" + Green)
}

func (txn Transaction) String() string {
	return fmt.Sprintf(Yellow+"'%v' should pay '%vDH' to '%v'", txn.Giver.Name, txn.Paid, txn.Receiver.Name+Yellow)
}

func New() (*Sharedb, error) {
	opts := badger.DefaultOptions("./badger")
	opts.Dir = "./Badger"
	opts.ValueDir = "./Badger"
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	sharedb := &Sharedb{db: db}
	return sharedb, nil
}

func (sh *Sharedb) Close() error {
	err := sh.db.Close()
	return err
}
