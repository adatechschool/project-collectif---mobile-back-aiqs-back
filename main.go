package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func UnmarshalSurf(data []byte) (Surf, error) {
	var r Surf
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Surf) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Surf struct {
	Records []Record `json:"records"`
}

type Record struct {
	ID     string `json:"id"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	SurfBreak               []string `json:"Surf Break"`
	DifficultyLevel         int64    `json:"Difficulty Level"`
	Destination             string   `json:"Destination"`
	MagicSeaweedLink        string   `json:"Magic Seaweed Link"`
	Photos                  []Photo  `json:"Photos"`
	PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`
	DestinationStateCountry string   `json:"Destination State/Country"`
	PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`
	Address                 string   `json:"Address"`
	Lat                     float64  `json:"lat"`
	Lng                     float64  `json:"lng"`
}

type Photo struct {
	ID         string     `json:"id"`
	Width      int64      `json:"width"`
	Height     int64      `json:"height"`
	URL        string     `json:"url"`
	Filename   string     `json:"filename"`
	Size       int64      `json:"size"`
	Type       string     `json:"type"`
	Thumbnails Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Small Full `json:"small"`
	Large Full `json:"large"`
	Full  Full `json:"full"`
}

type Full struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

	dsn := "root:root@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                        // default size for string fields
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // auto configure based on currently MySQL version
	}), &gorm.Config{})
}
