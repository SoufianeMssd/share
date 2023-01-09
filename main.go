package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"share/event"
	"time"
)

var participant1 = &event.Participant{
	Name:      "Soufiane MOUSSAID",
	Paid:      40,
	CreatedAt: time.Now(),
}

var participant2 = &event.Participant{
	Name:      "Ahmded Ali",
	Paid:      60,
	CreatedAt: time.Now(),
}

var participant3 = &event.Participant{
	Name:      "Arbi Benimellal",
	Paid:      20,
	CreatedAt: time.Now(),
}

var participant4 = &event.Participant{
	Name:      "Allal brahim",
	Paid:      30,
	CreatedAt: time.Now(),
}

var participant5 = &event.Participant{
	Name:      "Mohsine ramtane",
	Paid:      10,
	CreatedAt: time.Now(),
}

func main() {
	sharedb, err := event.New()
	handleErr(err)
	defer sharedb.Close()
	/* var pts []*event.Participant
	pts = append(pts, participant1, participant2, participant3, participant4, participant5) avg = 162.525

	event.SplitBill(pts) */
	action := flag.String("a", "list", "actions in share lib")
	help := flag.Bool("h", false, "help understand share lib")
	paid := flag.Float64("p", 0, "How much participant paid")
	flag.Parse()

	if *help {
		fmt.Println("this is help, will be configured later on")
	} else {
		switch *action {
		case "list":
			handleGetAll(sharedb)
		case "init":
			handleInitEvent(sharedb, flag.Args())
		case "get":
			handleFindEvent(sharedb, flag.Args())
		case "add":
			handleAddParticipant(sharedb, flag.Args(), *paid)
		case "remove":
			handleRemoveEvent(sharedb, flag.Args())
		case "split":
			handleSplitEventBill(sharedb, flag.Args())
		}
	}
}

func handleSplitEventBill(shdb *event.Sharedb, args []string) {
	evt := args[0]
	pts, err := shdb.Get(evt)
	handleErr(err)
	event.SplitBill(pts)
}

func handleRemoveEvent(shdb *event.Sharedb, args []string) {
	evt := args[0]
	err := shdb.Remove(evt)
	handleErr(err)
	fmt.Printf("Event '%v' removed succesfully\n", evt)
}

func handleAddParticipant(shdb *event.Sharedb, args []string, paid float64) {
	evt := args[0]
	pts_name := args[1]

	participant := event.Participant{
		Name:      pts_name,
		Paid:      paid,
		CreatedAt: time.Now(),
	}
	err := shdb.Add(evt, participant)
	handleErr(err)
	fmt.Printf("'%v' was added succesfully\n", pts_name)
}

func handleFindEvent(shdb *event.Sharedb, args []string) {
	evt := args[0]
	pts, err := shdb.Get(evt)
	handleErr(err)
	event.PrintPtsHeader()
	for _, it := range pts {
		fmt.Println(it)
	}
}

func handleInitEvent(shdb *event.Sharedb, args []string) {
	event := args[0]
	shdb.AddEvent(event)
	fmt.Printf("Event '%v' was added succesfully\n", event)
}

func handleGetAll(shdb *event.Sharedb) {
	events, err := shdb.GetAll()
	handleErr(err)
	if len(events) == 0 {
		fmt.Println("There are no events to list")
		os.Exit(0)
	}
	for _, evt := range events {
		fmt.Printf(event.Green+"Event : %v\n", evt.Name+event.Green)
		event.PrintParticipants(evt.Participants)
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
