package controllers

import (
    "encoding/json"
    "net/http"
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

	adddedData := map[string]string{
		"status": "Favorite added successfully",
	}

	// Encode and send JSON response
	json.NewEncoder(w).Encode(adddedData)
}

func handleGetFavorites(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    favorites := map[string][]string{
		"colors": {"red", "blue", "green"},
		"foods":  {"pizza", "sushi", "tacos"},
		"movies": {"Inception", "The Matrix", "Interstellar"},
	}
    
    // Encode and send JSON response
    json.NewEncoder(w).Encode(favorites)
}

func handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deletedData := map[string]string{
		"status": "Favorite deleted successfully",
	}

	// Encode and send JSON response
	json.NewEncoder(w).Encode(deletedData)
}

func handleUpdateFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	updatedData := map[string]string{
		"status": "Favorite updated successfully",
	}
	// Encode and send JSON response
	json.NewEncoder(w).Encode(updatedData)
}