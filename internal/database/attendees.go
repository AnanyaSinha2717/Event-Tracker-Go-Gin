package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"Id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}
