package model

import "time"

type Diary struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
}
