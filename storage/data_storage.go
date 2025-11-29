package storage

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"encoding/json"
	"os"
)

type DataStore struct {
	Charts []*models.Chart `json:"charts"`
	Insights []*models.Insight `json:"insights"`
	Audiences []*models.Audience `json:"audiences"`
	Users []models.User `json:"users"`
	Favorites []models.Favorite `json:"favorites"`
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