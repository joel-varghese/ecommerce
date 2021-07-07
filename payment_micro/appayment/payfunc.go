package appayment

import (
	"payments"
	"net/http"
)


func (a *App) MakePayment(w http.ResponseWriter, r *http.Request) {
	payments.MakePayment(a.DB, w, r)
}