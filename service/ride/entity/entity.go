package entity

import (
	"time"
)

type Ride struct {
	ID          string
	PassengerID string
	DriverID    string
	From        Location
	To          Location
	StartedAt   time.Time
	EndedAt     time.Time
}

type Location struct {
	Lat float64
	Lng float64
}
