package approduct

import (
	"products"
	"net/http"
)

func (a *App) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products.GetAllProducts(a.DB, w, r)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	products.CreateProduct(a.DB, w, r)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	products.GetProduct(a.DB, w, r)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products.GetProducts(a.DB, w, r)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	products.UpdateProduct(a.DB, w, r)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	products.DeleteProduct(a.DB, w, r)
}

func (a *App) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	products.GetProductByCategory(a.DB, w, r)
}