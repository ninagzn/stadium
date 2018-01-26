package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func ReservationPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("reservation-page.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func ReservationRequested(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	timeLayout := "15:04"
	date, _ := time.Parse("2006-01-02", r.Form.Get("date"))
	startTime, _ := time.Parse(timeLayout, r.Form.Get("startTime"))
	endTime, _ := time.Parse(timeLayout, r.Form.Get("endTime"))
	email := r.Form.Get("email")
	phoneNumber := r.Form.Get("phoneNumber")

	reservation := Reservation{
		Date:              date,
		StartTime:         startTime,
		EndTime:           endTime,
		ClientEmail:       email,
		ClientPhoneNumber: phoneNumber}

	t, err := template.ParseFiles("reservation-request-confirmation.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, reservation)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
