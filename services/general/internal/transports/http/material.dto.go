package http

import (
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
)

type Material struct {
	ID            uuid.UUID           `json:"id"`
	Kind          domain.MaterialKind `json:"kind"`
	Group         string              `json:"group"`
	Title         string              `json:"title"`
	Content       string              `json:"content"`
	Dependencies  []uuid.UUID         `json:"dependencies"`
	RequiredScore int64               `json:"requiredScore"`
	RequiredPeers int64               `json:"requiredPeers"`
	IsOptional    bool                `json:"isOptional"`
}

type CreateMaterial struct {
	Kind          domain.MaterialKind `json:"kind"`
	Group         string              `json:"group"`
	Title         string              `json:"title"`
	Content       string              `json:"content"`
	RequiredScore int64               `json:"requiredScore"`
	RequiredPeers int64               `json:"requiredPeers"`
	IsOptional    bool                `json:"isOptional"`
}

type UpdateMaterial struct {
	Group         optional.Option[string] `json:"group,omitempty"`
	Title         optional.Option[string] `json:"title,omitempty"`
	Content       optional.Option[string] `json:"content,omitempty"`
	RequiredScore optional.Option[int64]  `json:"requiredScore,omitempty"`
	RequiredPeers optional.Option[int64]  `json:"requiredPeers,omitempty"`
	IsOptional    optional.Option[bool]   `json:"isOptional,omitempty"`
}
