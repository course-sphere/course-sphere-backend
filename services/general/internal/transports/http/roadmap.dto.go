package http

import "github.com/google/uuid"

type CreateRoadmapData struct {
	Title       string `json:"title" description:"Roadmap title" example:"Backend Developer Roadmap"`
	Description string `json:"description" description:"Roadmap description" example:"A step-by-step roadmap for backend development"`
}

type AddCourseData struct {
	CourseID uuid.UUID `json:"course_id" format:"uuid" description:"ID of the course to add to the roadmap" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}

type Roadmap struct {
	AuthorID    uuid.UUID   `json:"author_id" format:"uuid" description:"ID of the roadmap author" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title       string      `json:"title" description:"Roadmap title" example:"Backend Developer Roadmap"`
	Description string      `json:"description" description:"Roadmap description" example:"A step-by-step roadmap for backend development"`
	Courses     []uuid.UUID `json:"courses,omitempty" format:"uuid" description:"List of course IDs in this roadmap"`
}

type MoveRoadmapCourseData struct {
	CurrentID uuid.UUID  `json:"current_id" format:"uuid" description:"UUID of the course to move" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	PrevID    *uuid.UUID `json:"prev_id,omitempty" format:"uuid" description:"UUID of the previous course" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	NextID    *uuid.UUID `json:"next_id,omitempty" format:"uuid" description:"UUID of the next course" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}

type UpdateRoadmapData struct {
	Title       *string `json:"title,omitempty" description:"Updated roadmap title"`
	Description *string `json:"description,omitempty" description:"Updated roadmap description"`
}
