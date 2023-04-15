package model

import "time"

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
