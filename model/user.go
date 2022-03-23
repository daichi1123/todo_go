package model

import "time"

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}
