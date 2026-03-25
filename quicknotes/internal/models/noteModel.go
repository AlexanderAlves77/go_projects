package models

import "time"

type Note struct {
	Id         int
	Title      string
	Content    string
	Author     string
	CreatedAt  time.Time
	UpdateddAt time.Time
}
