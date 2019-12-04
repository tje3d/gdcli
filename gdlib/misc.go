package gdlib

import uuid "github.com/satori/go.uuid"

// GenerateUUID generate a new uuid
func GenerateUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
