package src

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Key struct {
	User        User
	SigningKey  []byte
	TokenString string
}

var Keys []Key
var mySigningKey []byte

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Error")
	}
	key := mySigningKey
	mySigningKey = nil
	return key, nil
}

func setKey(tokenString string) {
	mySigningKey = nil
	for _, key := range Keys {
		if key.TokenString == tokenString {
			mySigningKey = key.SigningKey
		}
	}
	return
}

func IsAuthorised(endpoint func(writer http.ResponseWriter, request *http.Request)) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] != nil {

			tokenString := request.Header["Token"][0]
			setKey(tokenString)
			token, err := jwt.Parse(tokenString, getKey)

			if err != nil {
				fmt.Fprintf(writer, err.Error())
			}

			if token.Valid {
				endpoint(writer, request)
			}

		} else {
			fmt.Fprintf(writer, "Not Authorised")
		}
	})
}

func GetTokenUser(tokenString string) (User, error) {
	for _, key := range Keys {
		if key.TokenString == tokenString {
			return key.User, nil
		}
	}
	return User{}, fmt.Errorf("Error: Bad token string")
}

//
func GenerateJWT(user User) (string, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// make claims on token
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// generate token string
	str := strconv.Itoa(int(user.Id)) + user.Email
	signingKey := []byte(str)
	tokenString, err := token.SignedString(signingKey)

	// bounce if there was an error
	if err != nil {
		err = fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	key := Key{User: user, SigningKey: signingKey, TokenString: tokenString}
	Keys = append(Keys, key)

	// send back what we got
	return tokenString, nil
}
