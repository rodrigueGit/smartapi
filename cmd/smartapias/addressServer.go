package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Address structure to be used in the address registry
type Address struct {
	Key     string    `json:"key"`
	IP      string    `json:"ip"`
	Updated time.Time `json:"updated"`
}

var list map[string]Address

func main() {

	list = make(map[string]Address)
	router := mux.NewRouter()
	router.HandleFunc("/address", GetAddress).Methods("GET")
	router.HandleFunc("/address", AddAddress).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//GetAddress provides the list of registered addresses
func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}

//AddAddress registers a new address into the list of addresses
func AddAddress(w http.ResponseWriter, r *http.Request) {
	var address Address

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&address)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	list[address.Key] = address
	log.Println(address.Key)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}
