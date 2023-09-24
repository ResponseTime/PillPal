package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Kaisen", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}
func setSchedule(c *gin.Context) {
	item := c.Param("item")
	cron := cron.New()
	cron.AddFunc("@every 1m", func() { fmt.Println("reminder for ", item) })
	cron.Start()
	st := struct {
		Item string
		Set  bool
	}{Item: item, Set: true}
	c.IndentedJSON(http.StatusOK, st)
}
func main() {
	router := gin.Default()
	router.GET("api/albums", getAlbums)
	router.GET("api/set/:item", setSchedule)
	router.Run("localhost:8080")

}

// func inspect(entry []cron.Entry) {
// 	for _, entry := range entry {
// 		fmt.Println(entry.Job)
// 	}
// }
