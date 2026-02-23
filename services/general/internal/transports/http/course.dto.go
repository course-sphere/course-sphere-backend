package http

import (
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
)

type Course struct {
	Id         uuid.UUID `json:"id"`
	Thumbnail  string    `json:"thumbnail"`
	Title      string    `json:"title"`
	Tags       []string  `json:"tags,omitempty"`
	Instructor string    `json:"instructor"`
	Price      float32   `json:"price"`
}

type CreateCourse struct {
	Thumbnail string   `json:"thumbnail"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	Price     float32  `json:"price"`
}

type UpdateCourse struct {
	Thumbnail optional.Option[string]   `json:"thumbnail,omitempty"`
	Title     optional.Option[string]   `json:"title,omitempty"`
	Tags      optional.Option[[]string] `json:"tags,omitempty"`
	Price     optional.Option[float32]  `json:"price,omitempty"`
}

type CourseProgress struct {
	Progress float32 `json:"progress"`
	IsPassed bool    `json:"isPassed"`
}
