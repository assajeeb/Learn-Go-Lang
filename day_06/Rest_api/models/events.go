package models

import (
	"errors"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserId      int64
	Location    string
}

func (E *Event) Save() error {

	query := `
INSERT INTO events (name, description, location, dateTime, user_id)
VALUES (?, ?, ?, ?, ?)
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(E.Name, E.Description, E.Location, E.Time, E.UserId)
	if err != nil {
		panic(err)
		return err
	}

	id, err := result.LastInsertId()
	E.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
		return nil, errors.New("Unable to query")
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserId)
		if err != nil {
			panic(err)
			return nil, errors.New("Unable to scan")
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(&id)

	var event Event

	err = row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {

	query := `
	UPDATE events
	SET name=?, description=?, dateTime=?,location=?
	WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&event.Name, &event.Description, &event.Time, &event.Location, &event.Id)

	return err

}

func (event Event) Delete() error {
	query := `
	DELETE FROM events WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&event.Id)

	return err
}
