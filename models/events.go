package models

import (
	"example/com/db"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Event is a struct that represents an event in the database
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Location    string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// var events = []Event{}

// Save is a method that saves an event to the database
func (e Event) Save() error {
	query := `
	INSERT INTO events (Name, Location, Description, DateTime, UserID)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Location, e.Description, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Location, &e.Description, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil

}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Location, &e.Description, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET Name = ?, Location = ?, Description = ?, DateTime = ?
	WHERE ID = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Location, event.Description, event.DateTime, event.ID)
	return err
}
