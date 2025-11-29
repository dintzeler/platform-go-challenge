package routes

import (
    "net/http"
    "github.com/dintzeler/platform-go-challenge/controllers"
)

func SetupFavoritesRoutes() {
    http.HandleFunc("/favorites", controllers.FavoritesHandler)
}