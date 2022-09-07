package utils

import "github.com/google/uuid"

func NewUUID() string {
	ret, _ := uuid.NewUUID()
	return ret.String()
}
