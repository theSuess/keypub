package service

import "github.com/google/uuid"

func generateID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
