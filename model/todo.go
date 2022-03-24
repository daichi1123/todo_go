package model

import "time"

type (
	Todo struct {
		ID        int
		Content   string
		UserId    int
		CreatedAt time.Time
	}
)
