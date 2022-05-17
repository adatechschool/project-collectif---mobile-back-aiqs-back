package main

import (
	"wave/app"
	"wave/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

/* var db *sql.DB
var err error */

// func UnmarshalSurf(data []byte) (Surf, error) {
// 	var r Surf
// 	err := json.Unmarshal(data, &r)
// 	return r, err
// }

// func (r *Surf) Marshal() ([]byte, error) {
// 	return json.Marshal(r)
// }

/* type Surf struct {
	gorm.Model
	//ID                      uint    `gorm:"primaryKey"`
	SurfBreak               string  `json:"SurfBreak" gorm:"<-"`
	DifficultyLevel         int64   `json:"DifficultyLevel" gorm:"<-"`
	Destination             string  `json:"Destination" gorm:"<-"`
	MagicSeaweedLink        string  `json:"MagicSeaweedLink" gorm:"<-"`
	Photos                  string  `json:"Photos" gorm:"<-"`
	PeakSurfSeasonBegins    string  `json:"PeakSurfSeasonBegins" gorm:"<-"`
	DestinationStateCountry string  `json:"DestinationStateCountry" gorm:"<-"`
	PeakSurfSeasonEnds      string  `json:"PeakSurfSeasonEnds" gorm:"<-"`
	Address                 string  `json:"Address" gorm:"<-"`
	Lat                     float64 `json:"Lat" gorm:"<-"`
	Lng                     float64 `json:"Lng" gorm:"<-"`
} */

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome home!")
// }

/* func createEvent(w http.ResponseWriter, r *http.Request) {
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
*/
/* func createPost(w http.ResponseWriter, r *http.Request) {
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
} */

// func main() {
// 	/* gorm.Open(mysql.Open("surfdatabase.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")

// 	} */
// dsn := "root:root@tcp(127.0.0.1:8889)/surfdatabase"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	panic("failed to connect database")
// }

// 	// Migrate the schema
// 	db.AutoMigrate(&Surf{})

// Create
/* 	db.Create(&Surf{
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
	Lng:                     -158.05120469999997},
) */
/* 	db.Create(&Surf{

	ID:                      "2",
	SurfBreak:               "Point Break",
	DifficultyLevel:         5,
	Destination:             "Supertubes",
	MagicSeaweedLink:        "https://magicseaweed.com/Jeffreys-Bay-J-Bay-Surf-Report/88/",
	Photos:                  "https://dl.airtable.com/e3QoP3cFSyykZJOvWGIy_cesar-couto-477018-unsplash%20(1).jpg?ts=1652704905&userId=usrVSPQAdslUijuds&cs=1d36a62f395cd1d8",
	PeakSurfSeasonBegins:    "2018-08-01",
	DestinationStateCountry: "Jeffreys Bay, South Africa",
	PeakSurfSeasonEnds:      "2018-10-09",
	Address:                 "Supertubes, Jeffreys Bay, South Africa",
	Lat:                     -34.031783,
	Lng:                     24.93159400000002,
}) */

// db.Create(&Surf{

// 	SurfBreak:               "Reef Break",
// 	DifficultyLevel:         3,
// 	Destination:             "Pasta Point",
// 	MagicSeaweedLink:        "https://magicseaweed.com/Maldives-Surf-Forecast/56/",
// 	Photos:                  "https://dl.airtable.com/o4SxpoNxTSC49F4P2Hlc_aleksandar-popovski-693255-unsplash.jpg?ts=1652777877&userId=usrVSPQAdslUijuds&cs=c731b4899c3ef796",
// 	PeakSurfSeasonBegins:    "2018-04-01",
// 	DestinationStateCountry: "Maldives",
// 	PeakSurfSeasonEnds:      "2018-05-31",
// 	Address:                 "Pasta Point, Maldives",
// 	Lat:                     4.317842,
// 	Lng:                     73.59173299999998,
// })

// Read
//var surf Surf
//db.First(&surf, 1)                 // find product with integer primary key
//db.First(&surf, "code = ?", "D42") // find product with code D42

// Update - update product's price to 200
//db.Model(&surf).Update("Price", 200)
// Update - update multiple fields
//db.Model(&surf).Updates(Surf{Price: 200, Code: "F42"}) // non-zero fields
//db.Model(&surf).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// Delete - delete product
//	db.Delete(&surf, 1)

/* 	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/surfdatabase")
   	if err != nil {
   		panic(err.Error())
   	}
   	defer db.Close() */

// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homeLink)
// 	/* 	router.HandleFunc("/api/createSpot", createEvent).Methods("POST")
// 	   	router.HandleFunc("/api/spots", getAllEvents).Methods("GET")
// 	   	router.HandleFunc("/api/spots/{id}", getOneEvent).Methods("GET")
// 	   	router.HandleFunc("/api/spots/{id}", updateEvent).Methods("PATCH")
// 	   	router.HandleFunc("/api/spots/{id}", deleteEvent).Methods("DELETE") */
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
