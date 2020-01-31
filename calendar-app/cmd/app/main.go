package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/memstorage"

	"gopkg.in/yaml.v2"
)

// Config struct for configuring app
type Config struct {
	HTTPListen string `yaml:"http_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
}

var (
	configPath string
	config     Config
	logFile    *os.File
)

func init() {
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic("Can not read config file!")
	}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	logFile, err = os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logFile)
	log.SetPrefix(config.LogLevel + ": ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	defer logFile.Close()

	calendar := calendar.New(memstorage.NewEventsLocalStorage())

	t, err := time.Parse(time.RFC822, "26 Jan 20 19:00 MSK")
	if err != nil {
		log.Fatalln(err)
	}

	t2, err := time.Parse(time.RFC822, "25 Feb 21 11:30 MSK")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := calendar.AddEvent("Do cooking", "cook dinner", t2); err != nil {
		log.Fatalln(err)
	}

	if _, err := calendar.AddEvent("Go to party", "hang out!", t); err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", *r)
		for _, v := range calendar.GetEvents() {
			fmt.Fprintf(w, "Event: %s\nDescription: %s\nDate: %s\n\n", v.Name, v.Description, v.EventDate.String())
		}
	})

	http.ListenAndServe(config.HTTPListen, nil)
}
