package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/dintzeler/platform-go-challenge/services"
	"strconv"
	"fmt"
	"github.com/dintzeler/platform-go-challenge/models"
)

type AddFavoriteRequest struct {
	AssetID   int    `json:"asset_id"`
	AssetType models.AssetType `json:"asset_type"`
}

type UpdateFavoriteRequest struct {
	AssetID   int    `json:"asset_id"`
	AssetType models.AssetType `json:"asset_type"`
	Description string `json:"description"`
}

func FavortitesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetFavorites(w, r)
	case http.MethodPost:
		handleAddFavorite(w, r)
	case http.MethodDelete:
		handleDeleteFavorite(w, r)
	case http.MethodPut:
		handleUpdateFavorite(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleAddFavorite(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid User-ID", http.StatusBadRequest)
		return
	}

	var addFavoriteRequest AddFavoriteRequest
	err = json.NewDecoder(r.Body).Decode(&addFavoriteRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	adddedData := services.AddFavorite(userID, addFavoriteRequest.AssetID, addFavoriteRequest.AssetType)

	// Encode and send JSON response
	json.NewEncoder(w).Encode(adddedData)
}

func handleGetFavorites(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid User-ID", http.StatusBadRequest)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    favorites := services.GetFavorites(userID)
	fmt.Println("Fetched favorites:", favorites)
    
    // Encode and send JSON response
    json.NewEncoder(w).Encode(favorites)
}

func handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid User-ID", http.StatusBadRequest)
		return
	}
	assetIDStr := r.URL.Query().Get("asset_id")
	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		http.Error(w, "Invalid Asset-ID", http.StatusBadRequest)
		return
	}
	assetTypeStr := r.URL.Query().Get("asset_type")
	assetType := models.AssetType(assetTypeStr)

	w.Header().Set("Content-Type", "application/json")

	deletedData := services.DeleteFavorite(userID, assetID, assetType)

	// Encode and send JSON response
	json.NewEncoder(w).Encode(deletedData)
}

func handleUpdateFavorite(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-ID")
	userID, err := strconv.Atoi(userIDStr)
	w.Header().Set("Content-Type", "application/json")

	var updateFavorite UpdateFavoriteRequest
	err = json.NewDecoder(r.Body).Decode(&updateFavorite)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedData := services.UpdateFavorite(userID, updateFavorite.AssetType, updateFavorite.AssetID, updateFavorite.Description)

	// Encode and send JSON response
	json.NewEncoder(w).Encode(updatedData)
}