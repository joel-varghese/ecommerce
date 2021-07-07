package payments

import (
	"database/sql"
	"encoding/json"
	"context"
	"net/http"
	"log"
	"strconv"
	"model"
	uuid "github.com/google/uuid"
)

func MakePayment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	payment := model.Payment{}
	ctx := context.Background()
	type Price struct {
		Price string
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var cartID string

	stmt, err := db.Prepare("INSERT INTO payment(id,username,price,user_id) VALUES(?,?,?,?);")
	if err != nil {
		panic(err)
	}

	rows, err := db.QueryContext(ctx, "select price from cart where userid=?", payment.User_Id)
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
	uuid, _ := uuid.NewUUID()
	res, err := stmt.Exec(uuid, payment.Username,allPrice,payment.User_Id)

	if err != nil && res != nil {
		panic(err)
	}
	err0 := db.QueryRowContext(ctx, "DELETE FROM cart WHERE userid=?;", payment.User_Id).Scan(&cartID)
	if err0 != nil  {
		panic(err)
	}
	respondJSON(w, http.StatusCreated, payment)
}