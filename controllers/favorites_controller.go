package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/dintzeler/platform-go-challenge/services"
	"strconv"
	"fmt"
)

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
	w.Header().Set("Content-Type", "application/json")

	adddedData := services.AddFavorite(1, nil)


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
	w.Header().Set("Content-Type", "application/json")

	deletedData := services.DeleteFavorite("userID", "assetID")

	// Encode and send JSON response
	json.NewEncoder(w).Encode(deletedData)
}

func handleUpdateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	updatedData := services.UpdateFavorite("userID", nil)

	// Encode and send JSON response
	json.NewEncoder(w).Encode(updatedData)
}