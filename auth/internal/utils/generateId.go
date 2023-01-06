package utils

import "github.com/segmentio/ksuid"

func GenerateID() ksuid.KSUID {
	id := ksuid.New()
	return id
}
