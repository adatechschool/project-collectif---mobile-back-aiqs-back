package app

import (
	"fmt"
	"log"
	"net/http"

	"project-collectif---mobile-back-aiqs-back/config"
	"project-collectif---mobile-back-aiqs-back/handler"
	"project-collectif---mobile-back-aiqs-back/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/api/spots", a.GetAllEmployees)
	a.Post("/api/createSpot", a.CreateEmployee)
	a.Get("/api/spots/{id}", a.GetEmployee)
	a.Put("/api/spots/{id}", a.UpdateEmployee)
	a.Delete("/api/spots/{id}", a.DeleteEmployee)

}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a *App) GetAllSurfSpots(w http.ResponseWriter, r *http.Request) {
	handler.GetAllSurfSpots(a.DB, w, r)
}

func (a *App) CreateSurfSpots(w http.ResponseWriter, r *http.Request) {
	handler.CreateSurfSpots(a.DB, w, r)
}

func (a *App) GetSurfSpot(w http.ResponseWriter, r *http.Request) {
	handler.GetSurfSpot(a.DB, w, r)
}

func (a *App) UpdateSurfSpot(w http.ResponseWriter, r *http.Request) {
	handler.UpdateSurfSpot(a.DB, w, r)
}

func (a *App) DeleteSurfSpot(w http.ResponseWriter, r *http.Request) {
	handler.DeleteSurfSpot(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))

}
