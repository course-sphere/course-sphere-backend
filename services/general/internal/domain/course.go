package domain

import "github.com/google/uuid"

type Course struct {
	Id         uuid.UUID
	Thumbnail  string
	Title      string
	Tags       []string
	Instructor string
	Rating     float32
	Reviews    int64
	Students   int64
	Price      float32
}
