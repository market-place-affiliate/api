package uniqe

import "github.com/gofrs/uuid"

func UUID() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}
