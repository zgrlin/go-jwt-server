package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt"
)

var SigningKey = []byte("jwtsecret")

func webinterface(w http.ResponseWriter, r *http.Request) {
	validToken, err := CreateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["wallet"] = "0x"
	claims["expired"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/token", webinterface)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}

func main() {
	fmt.Println("JWT Web interface is run..")
	handleRequests()
}
