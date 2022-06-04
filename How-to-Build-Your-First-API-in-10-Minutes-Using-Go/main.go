package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{
		Name: "John Smith",
		Age:  30,
	},
	{
		Name: "Jane Doe",
		Age:  20,
	},
}

func main() {
	http.HandleFunc("/", serverStatus)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/create", usersCreateHandler)
	log.Printf("Listening...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serverStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: %s\n", http.StatusText(http.StatusOK))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func usersCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}
