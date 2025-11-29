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

func deleteFavorite(data *storage.DataStore, userID int, assetID int, assetType models.AssetType) error {
	var updatedFavorites []models.Favorite
	for _, fav := range data.Favorites {
		if !(fav.UserID == userID && fav.AssetID == assetID && fav.AssetType == assetType) {
			updatedFavorites = append(updatedFavorites, fav)
		}
	}

	data.Favorites = updatedFavorites
	err := storage.SaveData("data.json", data)
	if err != nil {
		return err
	}
	return nil
}


func updateFavorite(data *storage.DataStore, userID int, assetID int, assetType models.AssetType, description string) {
	assetExists := checkAssetExists(assetID, assetType)
	if !assetExists {
		fmt.Println("Asset does not exist")
		return
	}

	facoriteExists := checkFavoriteExists(getUserFavorites(userID), assetID, assetType)
	if !facoriteExists {
		fmt.Println("Favorite does not exist")
		return
	}

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

	storage.SaveData("data.json", data)


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

func DeleteFavorite(userID int, assetID int, assetType models.AssetType) map[string]string {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)

		return map[string]string{
			"status": "Error loading data",
		}
	}

	err = deleteFavorite(data, userID, assetID, assetType)
	if err != nil {
		return map[string]string{
			"status": "Error deleting favorite",
		}
	}
	
	return map[string]string{
		"status": "Favorite deleted successfully",
	}
}

func UpdateFavorite(userID int, assetType models.AssetType, assetID int, description string) map[string]string {
	data, err := storage.LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return map[string]string{
			"status": "Error loading data",
		}
	}

	updateFavorite(data, userID, assetID, assetType, description)

	return map[string]string{
		"status": "Favorite updated successfully",
	}
}