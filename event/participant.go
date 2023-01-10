package event

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

const Yellow = "\033[33m"
const Green = "\033[32m"

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
		fmt.Println(err)
		return nil, err
	}
	sharedb := &Sharedb{db: db}
	return sharedb, nil
}

func (sh *Sharedb) Close() error {
	err := sh.db.Close()
	return err
}
