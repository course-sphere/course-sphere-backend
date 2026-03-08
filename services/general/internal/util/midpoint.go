package util

import (
	"context"

	"github.com/google/uuid"
)

const (
	gap float64 = 1000
)

func Midpoint(
	ctx context.Context,
	id uuid.UUID,
	prevID *uuid.UUID,
	nextID *uuid.UUID,
	positionOf func(context.Context, uuid.UUID) (float64, error),
) (float64, error) {
	switch {
	case prevID == nil && nextID == nil:
		return gap, nil
	case prevID == nil:
		next, err := positionOf(ctx, *nextID)
		if err != nil {
			return 0, err
		}
		return next / 2, nil
	case nextID == nil:
		prev, err := positionOf(ctx, *prevID)
		if err != nil {
			return 0, err
		}
		return prev + gap, nil
	default:
		next, err := positionOf(ctx, *nextID)
		if err != nil {
			return 0, err
		}
		prev, err := positionOf(ctx, *prevID)
		if err != nil {
			return 0, err
		}

		return (prev + next) / 2, nil
	}
}
