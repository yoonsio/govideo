package models

import "time"

// Media -
type Media struct {
	ID    string
	Name  string
	Path  string
	Size  int64
	Added time.Time
}

// NewMedia -
func NewMedia() *Media {
	return &Media{}
}
