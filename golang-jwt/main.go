package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	fmt.Println("Generating a JWT...")
	token, err := generateJWT()
	if err != nil {
		log.Fatal("generateJWT: error generating JWT", err)
	}
	fmt.Println("JWT:", token)

	http.HandleFunc("/home", handlePage)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: error listening on port 8080", err)
	}
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		message = Message{"error", "Invalid request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		message = Message{"error", "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["authorized"] = true
	claims["user"] = "Hello.Kitty"
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
