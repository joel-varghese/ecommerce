package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"userid"`
	Quantity  int    `json:"quantity"`
	Price     float64    `json:"-"`
	Product_ID  uuid.UUID `json:"productid"`
}

type Product struct {
	ID          uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()";json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
}

type User struct {
	Id       uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Cart     string    `json:"cart"`
	Type     string    `json:"type"`

}

type Payment struct {
	Id       int       `json:"-"`
	Username string    `json:username`
	Price    float64   `json:price`
	User_Id  uuid.UUID `json:user_id`
}
type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}