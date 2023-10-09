package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var entries = make(map[string]cron.EntryID)
var cronI *cron.Cron

func main() {
	cronI = cron.New()
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/api/main", dataHandler)
	router.POST("api/set/", setSchedule)
	router.DELETE("api/del/", delSchedule)
	router.PATCH("/api/edit", editSchedule)
	router.POST("api/test", testRoute)
	router.Run("localhost:8080")

}
func editSchedule(c *gin.Context) {}

func dataHandler(c *gin.Context) {}

func setSchedule(c *gin.Context) {
	var dataNeeded struct {
		Medication string `json:"medication"`
		Time       string `json:"time"`
		Frequency  string `json:"frequency"`
	}
	if err := c.ShouldBindJSON(&dataNeeded); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedTime, err := time.Parse(time.RFC3339, dataNeeded.Time)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	switch strings.ToLower(dataNeeded.Frequency) {
	case "once":
		{
		}
	case "daily":
		{
		}
	case "weekly":
		{
		}
	}
	id, _ := cronI.AddFunc("@every 1m", func() { fmt.Println("reminder for ", dataNeeded) })
	fmt.Println(id)
	cronI.Start()
	st := struct {
		Item struct {
			Medication string `json:"medication"`
			Time       string `json:"time"`
			Frequency  string `json:"frequency"`
		}
		Set bool
	}{Item: dataNeeded, Set: true}
	c.IndentedJSON(http.StatusOK, st)
}

func delSchedule(c *gin.Context) {
	item := c.Param("item")
	cronI.Remove(entries[item])
	c.IndentedJSON(http.StatusOK, "done")
}

func testRoute(c *gin.Context) {
	var data struct {
		Param1 string `json:"param1"`
		Param2 int    `json:"param2"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param1 := data.Param1
	param2 := data.Param2
	c.JSON(http.StatusOK, gin.H{
		"param1": param1,
		"param2": param2,
	})
}
