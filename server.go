package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

func main() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/", HandlerDefault)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Server!\n")
	switch r.Method {
	case "GET":
		u := User{
			User:     "aditya",
			Password: "aditya1234",
		}
		fmt.Fprintf(w, "Incoming Get Request!\n")
		json.NewEncoder(w).Encode(u)
	case "POST":
		u := User{}
		fmt.Fprintf(w, "Incoming Post Request!\n")
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Fprintf(w, "User = %s\n", u.User)
		fmt.Fprintf(w, "Password = %s\n", u.Password)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
