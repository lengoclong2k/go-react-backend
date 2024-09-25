package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	MPAARating  string    `json:"mpaa_ratting"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdateAt    time.Time `json:"-"`
}
