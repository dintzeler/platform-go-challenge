package main

import (
	"fmt"
	"net/http"
    "github.com/dintzeler/platform-go-challenge/routes"
)





func main() {
	routes.SetupFavoritesRoutes()

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", nil)

}