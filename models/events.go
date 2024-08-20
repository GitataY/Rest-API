package models

import "time"

// Event is a struct that represents an event in the database
type Event struct {
	ID          int
	Name        string
	Location    string
	Description string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

// Save is a method that saves an event to the database
func (e Event) Save() {
	events = append(events, e)

} 

func GetAllEvents() []Event {
	return events

}