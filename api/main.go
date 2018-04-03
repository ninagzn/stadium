package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var reservations []Reservation

func GetReservations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reservations)
}

func GetReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range reservations {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Reservation{})
}

func Book(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reservation Reservation
	_ = json.NewDecoder(r.Body).Decode(&reservation)
	reservation.ID = params["id"]
	reservations = append(reservations, reservation)
	json.NewEncoder(w).Encode(reservations)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range reservations {
		if item.ID == params["id"] {
			reservations = append(reservations[:index], reservations[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(reservations)
	}
}

func main() {
	reservations = append(reservations, Reservation{ID: "1", ClientEmail: "client1@mail.com"})
	reservations = append(reservations, Reservation{ID: "2", ClientEmail: "client2@mail.com"})

	router := mux.NewRouter()
	router.HandleFunc("/reservations", GetReservations).Methods("GET")
	router.HandleFunc("/reservations/{id}", GetReservation).Methods("GET")
	router.HandleFunc("/book/{id}", Book).Methods("POST")
	router.HandleFunc("/cancel/{id}", CancelReservation).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
