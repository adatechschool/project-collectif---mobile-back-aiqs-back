package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func UnmarshalSurf(data []byte) (Surf, error) {
	var r Surf
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Surf) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Surf struct {
	ID                      string  `json:"id"`
	SurfBreak               string  `json:"SurfBreak"`
	DifficultyLevel         int64   `json:"DifficultyLevel"`
	Destination             string  `json:"Destination"`
	MagicSeaweedLink        string  `json:"MagicSeaweedLink"`
	Photos                  string  `json:"Photos"`
	PeakSurfSeasonBegins    string  `json:"PeakSurfSeasonBegins"`
	DestinationStateCountry string  `json:"DestinationStateCountry"`
	PeakSurfSeasonEnds      string  `json:"PeakSurfSeasonEnds"`
	Address                 string  `json:"Address"`
	Lat                     float64 `json:"lat"`
	Lng                     float64 `json:"lng"`
}

type allSurfSpots []Surf

var surfSpots = allSurfSpots{
	{
		ID:                      "1",
		SurfBreak:               "Reef Break",
		DifficultyLevel:         4,
		Destination:             "Pipeline",
		MagicSeaweedLink:        "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
		Photos:                  "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg?ts=1652704905&userId=usrVSPQAdslUijuds&cs=f1f4f021755f3333",
		PeakSurfSeasonBegins:    "2018-07-22",
		DestinationStateCountry: "Oahu, Hawaii",
		PeakSurfSeasonEnds:      "2018-08-31",
		Address:                 "Pipeline, Oahu, Hawaii",
		Lat:                     21.6650562,
		Lng:                     -158.05120469999997,
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newSpot Surf
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newSpot)
	surfSpots = append(surfSpots, newSpot)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpot)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range surfSpots {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(surfSpots)
}

/* func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var spots []Surf
	res, err := db.Query("SELECT * from surfspots")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	for res.Next() {
		var spot Surf
		err := res.Scan(&spot.ID, &spot.Destination)
		if err != nil {
			panic(err.Error())
		}
		spots = append(spots, spot)
	}
	json.NewEncoder(w).Encode(spots)
} */

/* func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var post Post
	  err := result.Scan(&post.ID, &post.Title)
	  if err != nil {
		panic(err.Error())
	  }
	  posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
  } */

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent Surf

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range surfSpots {
		if singleEvent.ID == eventID {
			singleEvent.SurfBreak = updatedEvent.SurfBreak
			singleEvent.DifficultyLevel = updatedEvent.DifficultyLevel
			singleEvent.Destination = updatedEvent.Destination
			singleEvent.MagicSeaweedLink = updatedEvent.MagicSeaweedLink
			singleEvent.Photos = updatedEvent.Photos
			singleEvent.PeakSurfSeasonBegins = updatedEvent.PeakSurfSeasonBegins
			singleEvent.PeakSurfSeasonEnds = updatedEvent.PeakSurfSeasonEnds
			singleEvent.Address = updatedEvent.Address
			singleEvent.Lat = updatedEvent.Lat
			singleEvent.Lng = updatedEvent.Lng

			surfSpots = append(surfSpots[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range surfSpots {
		if singleEvent.ID == eventID {
			surfSpots = append(surfSpots[:i], surfSpots[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/surfdatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/createSpot", createEvent).Methods("POST")
	router.HandleFunc("/api/spots", getAllEvents).Methods("GET")
	router.HandleFunc("/api/spots/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/api/spots/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/api/spots/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
