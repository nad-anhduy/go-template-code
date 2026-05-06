package util

import "github.com/google/uuid"

func GetUUID() string {

	id, err := uuid.NewV7()
	if err != nil {
		return uuid.New().String()
	}

	return id.String()
}
