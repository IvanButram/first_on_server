package models

import "time"

type ReadModel struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
