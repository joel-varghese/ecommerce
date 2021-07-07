package appuser

import (
	"user"
	"net/http"
)

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	user.GetAllUsers(a.DB, w, r)
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	user.RegisterUser(a.DB, w, r)
}

func (a *App) findUser(w http.ResponseWriter, r *http.Request) {
	user.FindUser(a.DB, w, r)
}

func (a *App) findUsers(w http.ResponseWriter, r *http.Request) {
	user.FindUsers(a.DB, w, r)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	user.UpdateUser(a.DB, w, r)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	user.DeleteUser(a.DB, w, r)
}