package main

import (
	"api/Router"
)

func main() {
	router := Router.SetupRouter()
	router.Run("localhost:8080")
}
