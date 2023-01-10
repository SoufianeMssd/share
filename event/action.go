package event

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sort"

	"github.com/dgraph-io/badger"
)

func (shdb *Sharedb) Add(event string, p Participant) error {
	pts, err := shdb.Get(event)
	if err != nil {
		return err
	}
	pts = append(pts, &p)

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(pts)

	return shdb.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(event), buffer.Bytes())
	})
}

func (shdb *Sharedb) AddEvent(event string) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode([]Participant{})
	return shdb.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(event), buffer.Bytes())
	})
}

func (shdb *Sharedb) Get(name string) ([]*Participant, error) {
	var pts []*Participant
	err := shdb.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err != nil {
			return err
		}
		pts, err = getRound(item)
		return err
	})

	return pts, err
}

func (shdb *Sharedb) GetAll() ([]Event, error) {
	var events []Event
	err := shdb.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			pts, err := getRound(item)
			if err != nil {
				return err
			}
			events = append(events, Event{Name: string(key), Participants: pts})
		}
		return nil
	})
	return events, err
}

func (shdb *Sharedb) Remove(event string) error {
	return shdb.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(event))
	})
}

func SplitBill(pts []*Participant) {
	sum := sum(pts)
	avg := sum / float64(len(pts))
	sortPts(pts)
	txns := getTxns(pts, sum, avg)
	PrintTxnsHeader()
	for _, txn := range txns {
		if txn.Paid > 0 {
			fmt.Println(txn)
		}
	}
}

func getTxns(pts []*Participant, sum float64, avg float64) []Transaction {
	var txns []Transaction

	for len(pts) > 1 {
		txn, err := getTxn(pts, avg)
		if err != nil {
			break
		}
		txns = append(txns, txn)
		pts = pts[1:]
		sortPts(pts)
	}
	return txns
}

func getTxn(pts []*Participant, avg float64) (Transaction, error) {
	var txn Transaction
	len := len(pts)

	diff := avg - pts[0].Paid
	if diff > 0 {
		pts[len-1].Paid = pts[len-1].Paid - diff
		txn.Giver = pts[0]
		txn.Receiver = pts[len-1]
		txn.Paid = diff
		return txn, nil
	}
	return txn, fmt.Errorf("Empty Transaction")
}

func sortPts(pts []*Participant) {
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].Paid < pts[j].Paid
	})
}

func sum(pts []*Participant) float64 {
	var sum float64 = 0
	for _, p := range pts {
		sum = sum + p.Paid
	}
	return sum
}

func getRound(item *badger.Item) ([]*Participant, error) {
	var buffer bytes.Buffer
	var pts []*Participant

	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})

	dec := gob.NewDecoder(&buffer)
	dec.Decode(&pts)
	return pts, err
}
