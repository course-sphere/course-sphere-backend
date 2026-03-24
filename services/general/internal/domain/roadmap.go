package domain

import "github.com/google/uuid"

type CreateRoadmapData struct {
	Title       string
	Description string
}

type Roadmap struct {
	AuthorID    uuid.UUID
	Title       string
	Description string
	Courses     []uuid.UUID
}

type UpdateRoadmapData struct {
	Title       *string
	Description *string
}
