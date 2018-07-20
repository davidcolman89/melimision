package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)


type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`

}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}


var people []Person


func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})


	router.HandleFunc("/people", GetPeople).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}



func GetPeople(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(people)
}
