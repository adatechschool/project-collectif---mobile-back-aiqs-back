package app

import (
	"log"
	"net/http"

	"wave/config"
	"wave/handler"
	"wave/model"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// "github.com/go-gorm/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dsn := "root:root@tcp(docker.for.mac.localhost:3306)/surfdatabase"
	//dsn := "docker:@tcp(docker.for.mac.localhost:3306)/surfdatabase"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Name,
	// 	config.DB.Charset)

	// db, err := gorm.Open(config.DB.Dialect, dbURI)
	// if err != nil {
	// 	log.Fatal("Could not connect database")
	// }

	a.DB = model.DBMigrate(db)
	a.DB = model.DBMigrateUser(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/api/users", a.GetAllUsers)
	a.Post("/api/createUser", a.CreateUser)
	a.Get("/api/spots", a.GetAllSurfSpots)
	a.Post("/api/createSpot", a.CreateSurfSpots)
	a.Get("/api/spots/{id}", a.GetSurfSpot)
	a.Put("/api/spots/{id}", a.UpdateSurfSpot)
	a.Delete("/api/spots/{id}", a.DeleteSurfSpot)

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

// Handlers to manage SurfSpots Data
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

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))

}
