package storage

import (
	"github.com/dintzeler/platform-go-challenge/models"
)

func GetById[T models.Asset](items []T, id int) T {
	for _, item := range items {
		if item.GetID() == id {
			return item
		}
	}

	var zero T
	return zero
}
