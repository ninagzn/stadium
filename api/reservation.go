package main

import "time"

type Reservation struct {
	ID                string
	Date              time.Time
	StartTime         time.Time
	EndTime           time.Time
	ClientEmail       string
	ClientPhoneNumber string
}
