package http

import "github.com/google/uuid"

type CreateRoadmapData struct {
	Title       string
	Description string
}

type AddCourseData struct {
	CourseID uuid.UUID
}

type Roadmap struct {
	AuthorID    uuid.UUID
	Position    float64
	Title       string
	Description string
	Courses     []uuid.UUID
}

type MoveRoadmapData struct {
	PrevID *uuid.UUID
	NextID *uuid.UUID
}

type UpdateRoadmapData struct {
	Position    *float64
	Title       *string
	Description *string
}
