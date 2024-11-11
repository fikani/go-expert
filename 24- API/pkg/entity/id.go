package entity

import "github.com/google/uuid"

// ID type
type ID = uuid.UUID

// NewID creates a new ID
func NewID() ID {
	return uuid.New()
}

// StringToID converts a string to an ID
func StringToID(s string) (ID, error) {
	return uuid.Parse(s)
}
