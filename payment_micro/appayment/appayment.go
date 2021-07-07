package appayment

import (
	"fmt"

	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	var err error

	a.DB, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/shopping_cart")

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

// Set all required routers
func (a *App) setRouters() {

	a.Post("/payment", a.MakePayment)

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

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
