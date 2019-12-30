package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	UUID string `json:"uuid"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UUID = uuid.New().String()
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.UUID == params["uuid"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})

}
func updateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.UUID == params["uuid"]{
			users = append(users[:index], users[index+1:]...)

			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UUID = params["uuid"]
			json.NewEncoder(w).Encode(&user)

			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.UUID == params["uuid"]{
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)

}

func main() {

	r := mux.NewRouter()

	users = append(users, User{
		UUID: uuid.New().String(),
		Firstname: "John",
		Lastname: "Bioux"})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{uuid}", getUser).Methods("GET")
	r.HandleFunc("/users/{uuid}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{uuid}", updateUser).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
 