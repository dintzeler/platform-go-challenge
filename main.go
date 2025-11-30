package main

import (
	"fmt"
	"net/http"
    "github.com/dintzeler/platform-go-challenge/routes"
	"github.com/joho/godotenv"
	"log"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	routes.SetupFavoritesRoutes()
	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", nil)

}