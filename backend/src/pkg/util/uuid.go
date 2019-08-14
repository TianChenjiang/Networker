package util

import (
	"github.com/google/uuid"
)

func MustUUID() string {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}

	return v
}

func NewUUID() (string, error)  {
	v, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return v.String(), nil
}
