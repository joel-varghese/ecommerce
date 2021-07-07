package carts

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"model"
	uuid "github.com/google/uuid"
)

func Getcarts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	type crt struct {
		UserID uuid.UUID `json:"userid"`
	}
	car := crt{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&car); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	rows, _ := db.QueryContext(ctx, "select * from cart where userid=?", car.UserID)
	carts := make([]model.Cart, 0)
	for rows.Next() {
		var cart model.Cart
		if err := rows.Scan( &cart.ID,&cart.UserID, &cart.Quantity,&cart.Price,&cart.Product_ID); err != nil {
			log.Fatal(err)
		}
		carts = append(carts, cart)
	}
	respondJSON(w, http.StatusOK, carts)
}

func AddToCart(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cart := model.Cart{}
	product := model.Product{}
	var price float64

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cart); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO cart(id,userid,quantity,price,product_id) VALUES(?,?,?,?,?);")

	if err != nil {
		panic(err)
	}
	prc, err := db.QueryContext(ctx, "select price from product where id=?", cart.Product_ID)
	for prc.Next() {
	if err := prc.Scan(&price); err != nil {
		log.Fatal(err)
		}
	}
	uuid, err := uuid.NewUUID()
	var carpice float64 = float64(cart.Quantity)
	res, err := stmt.Exec(uuid, cart.UserID, cart.Quantity, carpice*price, cart.Product_ID)

	if err != nil && res != nil {
		panic(err)
	}

	

	err0 := db.QueryRowContext(ctx, "select quantity from product WHERE id=?;",&cart.Product_ID).Scan(&product.Quantity)
	if err0 != nil {
		fmt.Println("not work")
		panic(err)
	}
	var quant int = product.Quantity - cart.Quantity
	err1 := db.QueryRowContext(ctx, "UPDATE product SET quantity=? WHERE id=?;",quant,cart.Product_ID).Scan(&product.Quantity)

	if err1 != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, cart)
}

func GetCartPrice(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	//cart := model.Cart{}

	type Cart struct {
		UserID uuid.UUID `json:"userid"`
	}

	type Price struct {
		Price string
	}

	cart := Cart{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cart); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select price from cart where userid=?", cart.UserID)
	
	if err != nil {
		panic(err)
	}

	cartPrice := make([]Price, 0)
	for rows.Next() {
		var price Price
		if err := rows.Scan(&price.Price); err != nil {
			log.Fatal(err)
		}

		cartPrice = append(cartPrice, price)
	}

	var allPrice int64
	for _, element := range cartPrice {

		i2, err := strconv.ParseInt(element.Price, 10, 64)

		if err != nil {
			panic(err)
		}

		allPrice += i2
	}
	respondJSON(w, http.StatusOK, allPrice)
}

func DeleteFromCart(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	type CartID struct {
		CartID uuid.UUID `json:"cartid"`
	}
	type ProdID struct {
		ProdID uuid.UUID `json:"prodid"`
	}

	cartID := CartID{}
	prodid := ProdID{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cartID); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	var quant int
	err0 := db.QueryRowContext(ctx, "select quantity,product_id from cart WHERE id=?;",cartID.CartID).Scan(&quant,&prodid.ProdID)
	if err0 != nil {
		fmt.Println("not work")
		panic(err0)
	}

	err := db.QueryRowContext(ctx, "DELETE FROM cart WHERE id=?;", cartID.CartID).Scan(&cartID.CartID)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		respondJSON(w, http.StatusOK, nil)
	}
	var Quantity string

	err1 := db.QueryRowContext(ctx, "UPDATE product SET quantity=quantity+? WHERE id=?;",quant,prodid.ProdID).Scan(&Quantity)

	if err1 != nil {
		panic(err)
	}

}