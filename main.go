package main

import (
	"fmt"
	"net/http"
    "github.com/dintzeler/platform-go-challenge/controllers"
)





func main() {
	http.HandleFunc("/favorites", controllers.FavortitesHandler)

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", nil)

}