package uuid

import "github.com/google/uuid"

func NewRequestId() string {
	return uuid.New().String()
}
