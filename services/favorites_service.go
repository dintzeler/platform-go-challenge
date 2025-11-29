package services

import (
	"github.com/dintzeler/platform-go-challenge/models"
)

func AddFavorite(userID string, asset models.Asset) map[string]string {
	return map[string]string{
		"status": "Favorite added successfully",
	}
}

func GetFavorites(userID string) map[string][]string {
	return map[string][]string{
		"colors": {"red", "blue", "green"},
		"foods":  {"pizza", "sushi", "tacos"},
		"movies": {"Inception", "The Matrix", "Interstellar"},
	}
}

func DeleteFavorite(userID string, assetID string) map[string]string {
	return map[string]string{
		"status": "Favorite deleted successfully",
	}
}

func UpdateFavorite(userID string, asset models.Asset) map[string]string {
	return map[string]string{
		"status": "Favorite updated successfully",
	}
}