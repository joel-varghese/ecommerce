package products

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"model"

	uuid "github.com/google/uuid"
)

func GetAllProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "select * from product")

	if err != nil {
		panic(err)
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	respondJSON(w, http.StatusOK, products)
}

func CreateProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO product(id,name,description,quantity,price) VALUES(?,?,?,?,?);")

	fmt.Println(db)
	if err != nil {
		fmt.Println("not work")
		panic(err)
	}

	uuid, err := uuid.NewUUID()

	res, err := stmt.Exec(uuid, product.Name,product.Description, product.Quantity, product.Price)

	if err != nil && res != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, product)
}

func GetProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}
	ctx := context.Background()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "SELECT * FROM product WHERE name=$1", product.Name).Scan(
		&product.ID, &product.Name,&product.Description, &product.Quantity, &product.Price)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:

		respondJSON(w, http.StatusOK, product)
	}
}

func GetProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select * from product where name=$1", product.Name)

	if err != nil {
		panic(err)
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name,&product.Description, &product.Quantity, &product.Price); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	respondJSON(w, http.StatusOK, products)
}

func UpdateProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "UPDATE product SET name=?,description=?, quantity=?, price=? WHERE name=?;",
		&product.Name,&product.Description, &product.Quantity, &product.Price, &product.Name).Scan(
		&product.Name,&product.Description, &product.Quantity, &product.Price, &product.Name)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, product)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		respondJSON(w, http.StatusOK, product)
	}
}

func DeleteProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "DELETE FROM product WHERE name=$1;", product.Name).Scan(&product.Name)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:

		respondJSON(w, http.StatusOK, nil)
	}
}

func GetProductByCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select * from product where name=$1", product.Name)

	if err != nil {
		panic(err)
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name,&product.Description, &product.Quantity, &product.Price); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	respondJSON(w, http.StatusOK, products)
}