package storage

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"os"
)

func GetChartById(charts []models.Chart, id int) *models.Chart {
	for _, chart := range charts {
		if chart.GetID() == id {
			return &chart
		}
	}
	return nil
}

func GetInsightById(insights []models.Insight, id int) *models.Insight {
	for _, insight := range insights {
		if insight.GetID() == id {
			return &insight
		}
	}
	return nil
}

func GetAudienceById(audiences []models.Audience, id int) *models.Audience {
	for _, audience := range audiences {
		if audience.GetID() == id {
			return &audience
		}
	}
	return nil
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
		asset := GetChartById(data.Charts, assetID)
		return asset != nil
	case models.AssetTypeInsight:
		asset := GetInsightById(data.Insights, assetID)
		return asset != nil
	case models.AssetTypeAudience:
		asset := GetAudienceById(data.Audiences, assetID)
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