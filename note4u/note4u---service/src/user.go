package src

import (
	"fmt"
	"net/http"
)

type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TimeCreated string `json:"createAt"`
	Notes       []Note
}

//
func (user *User) Create() (err error) {
	statement := "insert into users (email, password) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return
	}
	return
}

func Users(id int) (user User, err error) {

	err = Db.QueryRow("select id, email, password, created_at from users where id=$1", id).Scan(&user.Id, &user.Email, &user.Password, &user.TimeCreated)
	if err != nil {
		return
	}

	rows, err := Db.Query("select id, title, body, user_id, created_at from notes where user_id=$1", id)
	if err != nil {
		return user, err
	}

	// this will be used to create the notes in order to break the connection
	// else, printing or marshal will cause a lot of problems, since one point to the other and the other points back
	temp_user := User{}
	user.Notes = []Note{}
	for rows.Next() {
		note := Note{}

		err = rows.Scan(&note.Id, &note.Title, &note.Body, &temp_user.Id, &note.TimeCreated)
		if err != nil {
			return user, err
		}

		user.Notes = append(user.Notes, note)
	}

	rows.Close()
	return
}

func (user *User) Update() (err error) {
	_, err = Db.Exec("update users set email=$2, password=$3 where id=$1", user.Id, user.Email, user.Password)
	return
}

func (user *User) Delete() (err error) {
	_, err = Db.Exec("delete from notes where user_id = $1", user.Id)
	if err != nil {
		return
	}

	_, err = Db.Exec("delete from users where id = $1", user.Id)
	return
}

//
func GetUser(email, password string) (user User, err error) {
	user = User{}

	err = Db.QueryRow(`SELECT id, email, password, timeCreated FROM users WHERE email = $1 AND password = $2;`, email, password).Scan(&user.Id, &user.Email, &user.Password, &user.TimeCreated)

	return
}

//
func GetRequesttingUser(request *http.Request) (user User, err error) {

	// get token string
	var tokenString string
	if request.Header["Token"] != nil {
		tokenString = request.Header["Token"][0]
	} else {
		err = fmt.Errorf("Error: token string was not provided")
		return
	}

	// use token string to get user
	user, err = GetTokenUser(tokenString)
	if err != nil {
		err = fmt.Errorf("Error: finding user, bad token string")
		return
	}

	return
}
