package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/memstorage"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	calendar := calendar.New(memstorage.NewEventsLocalStorage())
	t, err := time.Parse(time.RFC822, "26 Jan 20 19:00 MSK")
	if err != nil {
		log.Fatalln(err)
	}

	if err := calendar.AddEvent("Go to party", "hang out!", t); err != nil {
		log.Fatalln(err)
	}

	t2, err := time.Parse(time.RFC822, "25 Feb 21 11:30 MSK")
	if err != nil {
		log.Fatalln(err)
	}

	if err := calendar.AddEvent("Go to university", "learning in university...", t2); err != nil {
		log.Fatalln(err)
	}

	for _, v := range calendar.GetEvents() {
		fmt.Printf("%+v\n", v)
	}
}
