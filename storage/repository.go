package storage

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"os"
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

func GetUserFavorites(data *DataStore, userID int) []models.Favorite {
	var userFavorites []models.Favorite
	for _, fav := range data.Favorites {
		if fav.UserID == userID {
			userFavorites = append(userFavorites, fav)
		}
	}
	return userFavorites
}

func CreateFavorite(data *DataStore, userID int, assetID int, assetType models.AssetType) error {
	newFavorite := models.Favorite{
		UserID: userID,
		AssetID: assetID,
		AssetType: assetType,
	}

    data.Favorites = append(data.Favorites, newFavorite)
    return SaveData(os.Getenv("DATA_FILE"), data)
}

func FavoriteExists(favorites []models.Favorite, assetID int, assetType models.AssetType) bool {
    for _, fav := range favorites {
        if fav.AssetID == assetID && fav.AssetType == assetType {
            return true
        }
    }
    return false
}

func AssetExists(data *DataStore,assetID int, assetType models.AssetType) bool {
	switch assetType {
	case models.AssetTypeChart:
		asset := GetById(data.Charts, assetID)
		return asset != nil
	case models.AssetTypeInsight:
		asset := GetById(data.Insights, assetID)
		return asset != nil
	case models.AssetTypeAudience:
		asset := GetById(data.Audiences, assetID)
		return asset != nil
	default:
		return false
	}
}

func UpdateAssetDescription(data *DataStore, userID int, assetID int, assetType models.AssetType, description string) error {
	switch assetType {
	case models.AssetTypeChart:
		for _, chart := range data.Charts {
			if chart.GetID() == assetID {
				chart.SetDescription(description)
			}
		}
	case models.AssetTypeInsight:
		for _, insight := range data.Insights {
			if insight.GetID() == assetID {
				insight.SetDescription(description)
			}
		}
	case models.AssetTypeAudience:
		for _, audience := range data.Audiences {
			if audience.GetID() == assetID {
				audience.SetDescription(description)
			}
		}
	}

	return SaveData(os.Getenv("DATA_FILE"), data)
}

func RemoveFavorite(data *DataStore, userID int, assetID int, assetType models.AssetType) error {
	var updatedFavorites []models.Favorite
	for _, fav := range data.Favorites {
		if !(fav.UserID == userID && fav.AssetID == assetID && fav.AssetType == assetType) {
			updatedFavorites = append(updatedFavorites, fav)
		}
	}

	data.Favorites = updatedFavorites
	return SaveData(os.Getenv("DATA_FILE"), data)
}