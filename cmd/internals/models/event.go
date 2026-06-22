package models

import "time"

type Event struct {
	ID         uint
	Timestamp  time.Time
	Endpoint   string
	Method     string
	StatusCode int
	LatencyMs  float64
}
