package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var users []User

var count int = 0

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", GetUserByID)
	r.HandleFunc("/user", GetUsers)
	http.ListenAndServe(":8000", r)
	fmt.Printf("Server listing at 8000")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count++
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = count
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "User Id doesn't exist")
	}
	for index, item := range users {
		if item.ID == userId {
			users = append(users[:index], users[:index+1]...)
			var user User

			_ = json.NewDecoder(r.Body).Decode(user)
			user.ID = userId
			users = append(users, user)
			json.NewEncoder(w).Encode(&user)
			return
		}
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "User Id doesn't exist")
	}
	for _, item := range users {
		if item.ID == userId {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
