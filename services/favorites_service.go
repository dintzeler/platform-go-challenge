package services

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"encoding/json"
	"os"
	"fmt"
)

type DataStore struct {
	Charts []models.Chart `json:"charts"`
	Insights []models.Insight `json:"insights"`
	Audiences []models.Audience `json:"audiences"`
	Users []models.User `json:"users"`
	Favorites []models.Favorite `json:"favorites"`
}

type FavoritesResponse struct {
	Charts	 []models.Chart `json:"charts"`
	Insights []models.Insight `json:"insights"`
	Audiences []models.Audience `json:"audiences"`
}

func LoadData(fileName string) (*DataStore, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var dataStore DataStore
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dataStore)
	if err != nil {
		return nil, err
	}
	return &dataStore, nil
}

func searchFavoritesByUserID(data *DataStore, userID int) []models.Favorite {
	var userFavorites []models.Favorite
	for _, fav := range data.Favorites {
		if fav.UserID == userID {
			userFavorites = append(userFavorites, fav)
		}
	}
	return userFavorites
}

func AddFavorite(userID int, asset models.Asset) map[string]string {
	return map[string]string{
		"status": "Favorite added successfully",
	}
}

func buildFavoritesResponse(favorites []models.Favorite, data *DataStore) FavoritesResponse {
	favoritesResponse := FavoritesResponse{
		Charts:   []models.Chart{},
		Insights: []models.Insight{},
		Audiences: []models.Audience{},
	}
	for _, fav := range favorites {
		switch fav.AssetType {
		case models.AssetTypeChart:
			for _, chart := range data.Charts {
				if chart.ID == fav.AssetID {
					favoritesResponse.Charts = append(favoritesResponse.Charts, chart)
				}
			}
		case models.AssetTypeInsight:
			for _, insight := range data.Insights {
				if insight.ID == fav.AssetID {
					favoritesResponse.Insights = append(favoritesResponse.Insights, insight)
				}
			}
		case models.AssetTypeAudience:
			for _, audience := range data.Audiences {
				if audience.ID == fav.AssetID {
					favoritesResponse.Audiences = append(favoritesResponse.Audiences, audience)
				}
			}
		}
	}
	return favoritesResponse
}



func GetFavorites(userID int) FavoritesResponse {
	data, err := LoadData("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return FavoritesResponse{}
	}
	
	favorites := searchFavoritesByUserID(data, userID)
	favoritesResponse := buildFavoritesResponse(favorites, data)
	
	return favoritesResponse
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