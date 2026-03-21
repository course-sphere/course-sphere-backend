package domain

import "github.com/google/uuid"

type CreateRoadmapData struct {
	Title       string
	Description string
}

type Roadmap struct {
	AuthorID    uuid.UUID
	Position    float64
	Title       string
	Description string
	Courses     []uuid.UUID
}

type UpdateRoadmapData struct {
	Position    *float64
	Title       *string
	Description *string
}
