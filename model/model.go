package model

import "time"

// Model represents meta data of entity.
type Model struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
