package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	uuid "github.com/google/uuid"

	"model"
)

func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
		rows, err := db.QueryContext(ctx, "select * from user")
		if err != nil {
			panic(err)
		}
		users := make([]model.User, 0)
		for rows.Next() {
			var user model.User
			if err := rows.Scan( &user.Id,&user.Name, &user.Cart,&user.Type); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		respondJSON(w, http.StatusOK, users)
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO user(id,name,cart,type) VALUES(?,?,?,?);")

	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewUUID()

	res, err := stmt.Exec(uuid, user.Name,user.Cart,user.Type)

	if err != nil && res != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, user)
}

func FindUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	ctx := context.Background()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "SELECT * FROM user WHERE name=$1", user.Name).Scan(
		&user.Id,&user.Name,&user.Cart,&user.Type)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created \n", user.Name)
		respondJSON(w, http.StatusOK, user)
	}
}

func FindUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select * from user where name=$1", user.Name)

	if err != nil {
		panic(err)
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan( &user.Id,&user.Name,&user.Cart,&user.Type); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	respondJSON(w, http.StatusOK, users)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "UPDATE user SET name=$1, cart=$2,type=$4 WHERE name=$5;",
		user.Name, user.Cart,user.Type, user.Name).Scan(
			&user.Id,&user.Name, &user.Cart,&user.Type)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, user)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created \n", user.Name)
		respondJSON(w, http.StatusOK, user)
	}
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "DELETE FROM user WHERE name=$1;", user.Name).Scan(&user.Name)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created \n", user.Name)
		respondJSON(w, http.StatusOK, nil)
	}

}

//LoginUser
/*
func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	type Credentials struct {
		Id       uuid.UUID `json:"-"`
		Mobile   string    `json:"mobile", db:"mobile"`
		Password string    `json:"password", db:"password"`
	}

	type UserToken struct {
		Token string `json:"token"`
	}

	user := Credentials{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	userPassword := user.Password

	err := db.QueryRowContext(ctx, "SELECT id, mobile, password FROM users WHERE mobile=$1", user.Mobile).Scan(
		&user.Id, &user.Mobile, &user.Password)

	switch {
	case err == sql.ErrNoRows:

		respondError(w, http.StatusNotFound, err.Error())
		return
	case err != nil:

		log.Fatalf("query error: %v\n", err)
	default:

		match := utils.CheckPasswordHash(userPassword, user.Password)

		if !match {

			respondError(w, http.StatusBadRequest, "wrong password")
			return
		}

		userToken := &UserToken{}
		token := utils.MakeTokenFromUUID(user.Id)
		userToken.Token = token
		respondJSON(w, http.StatusOK, userToken)

	}

}

func SendSmsVerfication(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	mobile := user.Mobile
	password := user.Password

	if len(mobile) == 11 && strings.HasPrefix(mobile, "09") {
		if len(password) >= 6 {
			respondJSON(w, http.StatusOK, "کد فعال سازی به شماره وارد شده ارسال شد")
			userCode := SendVerficationSms(mobile)

			stmt, err := db.Prepare("INSERT INTO user_auth(mobile, password, code) VALUES($1,$2,$3);")

			if err != nil {
				panic(err)
			}

			userPass, _ := utils.HashPassword(user.Password)

			res, err := stmt.Exec(user.Mobile, userPass, userCode)

			if err != nil && res != nil {
				panic(err)
			}

			respondJSON(w, http.StatusCreated, user)

		} else {
			respondError(w, http.StatusNotFound, "رمز عبور حداقل باید ۶ کاراکتر باشد")
		}
	} else {
		respondError(w, http.StatusNotFound, "لطفا شماره موبایل را به درستی وارد نمایید")
	}

}

func GetOtpFromUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	type UserData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}

	userData := UserData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	err := db.QueryRowContext(ctx, "SELECT * FROM user_auth WHERE mobile=$1 And code=$2",
		userData.Mobile, userData.Code).Scan(
		&userData.Mobile, &userData.Password, &userData.Code)

	switch {
	case err == sql.ErrNoRows:
		respondError(w, http.StatusNotFound, err.Error())
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:

		uuid, _ := uuid.NewUUID()

		stmt, err := db.Prepare("INSERT INTO users(name,avatar,mobile,birth_day,identify,cart,credit,password,type,email,id) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);")

		if err != nil {
			panic(err)
		}

		res, err := stmt.Exec("", "", userData.Mobile, "", "", "", "", userData.Password, "", "", uuid)

		if err != nil && res != nil {
			panic(err)
		}
		respondJSON(w, http.StatusOK, userData)
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	type UserToken struct {
		Token string `json:"token"`
	}

	reqToken := r.Header.Get("Authorization")

	userID, okExtractUUID := utils.ExtractClaimsForRefresh(reqToken)

	if !okExtractUUID {
		fmt.Println("Can't get UUID from token in refresh section")
		return
	}

	userUUID, errUUID := uuid.Parse(userID["id"].(string))
	if errUUID != nil {
		panic(errUUID)
	}

	token := utils.MakeTokenFromUUID(userUUID)

	userToken := &UserToken{}
	userToken.Token = token
	respondJSON(w, http.StatusOK, userToken)
}
*/