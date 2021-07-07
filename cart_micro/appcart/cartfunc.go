package appcart

import (
	"carts"
	"net/http"
)


func (a *App) Getcarts(w http.ResponseWriter, r *http.Request) {
	carts.Getcarts(a.DB, w, r)
}

func (a *App) AddToCart(w http.ResponseWriter, r *http.Request) {
	carts.AddToCart(a.DB, w, r)
}

func (a *App) GetCartPrice(w http.ResponseWriter, r *http.Request) {
	carts.GetCartPrice(a.DB, w, r)
}

func (a *App) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	carts.DeleteFromCart(a.DB, w, r)
}