package http

import "github.com/google/uuid"

type Course struct {
    Id         uuid.UUID `json:"id"`
    Thumbnail  string    `json:"thumbnail"`
    Title      string    `json:"title"`
    Tags       []string  `json:"tags,omitempty"`
    Instructor string    `json:"instructor"`
    Rating     float32   `json:"rating"`
    Reviews    int64     `json:"reviews"`
    Students   int64     `json:"students"`
    Price      float32   `json:"price"`
}
