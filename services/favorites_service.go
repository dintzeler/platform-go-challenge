package services

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"fmt"
	"github.com/dintzeler/platform-go-challenge/storage"
)

type FavoritesResponse struct {
	Charts	 []models.Chart `json:"charts"`
	Insights []models.Insight `json:"insights"`
	Audiences []models.Audience `json:"audiences"`
}


func searchFavoritesByUserID(data *storage.DataStore, userID int) []models.Favorite {
	var userFavorites []models.Favorite
	for _, fav := range data.Favorites {
		if fav.UserID == userID {
			userFavorites = append(userFavorites, fav)
		}
	}
	return userFavorites
}


func getUserFavorites(userID int) []models.Favorite {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return []models.Favorite{}
	}

	return searchFavoritesByUserID(data, userID)
}

func checkFavoriteExists(favorites []models.Favorite, assetID int, assetType models.AssetType) bool {
	for _, fav := range favorites {
		if fav.AssetID == assetID && fav.AssetType == assetType {
			return true
		}
	}
	return false
}

func addFavorite(userID int, assetID int, assetType models.AssetType) error {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return err
	}

	newFavorite := models.Favorite{
		UserID:    userID,
		AssetID:   assetID,
		AssetType: assetType,
	}

	data.Favorites = append(data.Favorites, newFavorite)
	err = storage.SaveData("data.json", data)
	if err != nil {
		return err
	}
	return nil
}

func checkAssetExists(assetID int, assetType models.AssetType) bool {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return false
	}
	switch assetType {
	case models.AssetTypeChart:
		asset := storage.GetById(data.Charts, assetID)
		return asset != nil
	case models.AssetTypeInsight:
		asset := storage.GetById(data.Insights, assetID)
		return asset != nil
	case models.AssetTypeAudience:
		asset := storage.GetById(data.Audiences, assetID)
		return asset != nil
	default:
		return false
	}
}

func buildFavoritesResponse(favorites []models.Favorite, data *storage.DataStore) FavoritesResponse {
	favoritesResponse := FavoritesResponse{
		Charts:   []models.Chart{},
		Insights: []models.Insight{},
		Audiences: []models.Audience{},
	}
    for _, fav := range favorites {
        switch fav.AssetType {
        case models.AssetTypeChart:
			chart := storage.GetById(data.Charts, fav.AssetID)
            if chart != nil {
                favoritesResponse.Charts = append(favoritesResponse.Charts, *chart)
            }
        case models.AssetTypeInsight:
			insight := storage.GetById(data.Insights, fav.AssetID)
            if insight != nil {
                favoritesResponse.Insights = append(favoritesResponse.Insights, *insight)
            }
        case models.AssetTypeAudience:
			audience := storage.GetById(data.Audiences, fav.AssetID)
            if audience != nil {
                favoritesResponse.Audiences = append(favoritesResponse.Audiences, *audience)
            }
        }
    }
	return favoritesResponse
}



func GetFavorites(userID int) FavoritesResponse {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return FavoritesResponse{}
	}
	
	favorites := searchFavoritesByUserID(data, userID)
	favoritesResponse := buildFavoritesResponse(favorites, data)
	
	return favoritesResponse
}

func AddFavorite(userID int, assetID int, assetType models.AssetType) map[string]string {
	userFavorites := getUserFavorites(userID)
	favoriteExists := checkFavoriteExists(userFavorites, assetID, assetType)
	if favoriteExists {
		return map[string]string{
			"status": "Favorite already exists",
		}
	}

	assetExists := checkAssetExists(assetID, assetType)
	if !assetExists {
		return map[string]string{
			"status": "Asset does not exist",
		}
	}

	err := addFavorite(userID, assetID, assetType)
	if err != nil {
		return map[string]string{
			"status": "Error adding favorite",
		}
	}

	return map[string]string{
		"status": "Favorite added successfully",
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