package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
  "io/ioutil"
)

//Delivery structure
type Delivery struct {
	FileName string
	Data     []byte
  Sent     time.Time
}

var list map[string]Delivery

func main() {
  list = make(map[string]Delivery)
	router := mux.NewRouter()
	router.HandleFunc("/file", GetFiles).Methods("GET")
  router.HandleFunc("/file", AddFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

//GetFiles provide the list of delivered files
func GetFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}

//AddFile add a new file into the list of delivery files
func AddFile(w http.ResponseWriter, r *http.Request) {
	var delivery Delivery

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&delivery)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	list[delivery.FileName] = delivery
	log.Println("File delivered:", delivery.FileName)

  err = ioutil.WriteFile(delivery.FileName, delivery.Data, 0644)
  if err!=nil {
    panic(err)
  }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		panic(err)
	}
}
