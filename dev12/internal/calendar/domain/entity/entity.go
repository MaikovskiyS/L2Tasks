package entity

import "time"

type Event struct {
	Id          string
	Name        string
	Date        time.Time
	Description string
}
