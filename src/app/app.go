package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jestape/hackovid-dyb-api/src/app/handler"
	"github.com/jestape/hackovid-dyb-api/src/app/model"
	"github.com/jestape/hackovid-dyb-api/src/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.Charset)

	db, err := gorm.Open(config.Database.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects

	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world! This is the DoYourBit api."))
	})

	a.Get("/users/{pk}", a.handleRequest(handler.GetUser))
	a.Get("/users", a.handleRequest(handler.GetUsers))
	a.Post("/seller", a.handleRequest(handler.CreateSeller))
	a.Post("/buyer", a.handleRequest(handler.CreateBuyer))

	a.Get("/products/{id}", a.handleRequest(handler.GetProduct))
	a.Get("/products", a.handleRequest(handler.GetProducts))
	a.Post("/products", a.handleRequest(handler.CreateProduct))

	a.Get("/tickets/{id}", a.handleRequest(handler.GetTicket))
	a.Get("/tickets", a.handleRequest(handler.GetTicketsUserB)).Queries("buyer", "{buyer}")
	a.Get("/tickets", a.handleRequest(handler.GetTicketsUserS)).Queries("seller", "{seller}")
	a.Get("/tickets", a.handleRequest(handler.GetTicketTxHash)).Queries("tx_hash", "{tx_hash}")
	a.Get("/tickets", a.handleRequest(handler.GetTickets))
	a.Post("/tickets", a.handleRequest(handler.CreateTicket))
	a.Put("/tickets/{id}", a.handleRequest(handler.UpdateTicket))

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowCredentials(),
	)

	a.Router.Use(cors)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("GET", "OPTIONS")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("POST", "OPTIONS")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("PUT", "OPTIONS")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("DELETE", "OPTIONS")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
