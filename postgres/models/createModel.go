package models

import "time"

type CreateModel struct {
	Title       string
	Description string
	CreatedAt   time.Time
}
