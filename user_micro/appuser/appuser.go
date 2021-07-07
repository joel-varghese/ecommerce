package appuser

import (
	"fmt"

	"log"
	"net/http" 

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var err error

	a.DB, err = sql.Open("mysql", "root:mysql@tcp(host.docker.internal:3306)/shopping_cart")

	if err != nil {
		log.Fatal("Could not connect database")
	}

	err = a.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	a.Router = mux.NewRouter()
	a.setRouters()

}

func (a *App) setRouters() {
	// Routing for handling the projects

	a.Get("/users", a.getAllUsers)
	a.Post("/user", a.registerUser)
	a.Get("/user/finduser", a.findUser)
	a.Get("/user/findusers", a.findUsers)
	a.Put("/user", a.updateUser)
	a.Delete("/user", a.deleteUser)

}

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

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}