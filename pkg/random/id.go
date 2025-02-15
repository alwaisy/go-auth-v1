package random

import "github.com/google/uuid"

func NewRandomID() uuid.UUID {
	return uuid.New()
}
