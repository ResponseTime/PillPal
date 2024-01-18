package Scheduler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var entries = make(map[string]cron.EntryID)
var cronI *cron.Cron = cron.New()

func EditSchedule(c *gin.Context) {}

func SetSchedule(c *gin.Context) {
	var dataNeeded struct {
		Medication string `json:"medication"`
		Time       string `json:"time"`
		Frequency  string `json:"frequency"`
	}
	if err := c.ShouldBindJSON(&dataNeeded); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := time.Parse(time.RFC3339, dataNeeded.Time)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	var id cron.EntryID
	switch strings.ToLower(dataNeeded.Frequency) {
	case "once":
		{
			id, _ = cronI.AddFunc("@every 1m", func() { fmt.Println("reminder for ", dataNeeded) })
		}
	case "daily":
		{
			id, _ = cronI.AddFunc("@every 1m", func() { fmt.Println("reminder for ", dataNeeded) })
		}
	case "weekly":
		{
			id, _ = cronI.AddFunc("@every 1m", func() { fmt.Println("reminder for ", dataNeeded) })
		}
	}

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

func DelSchedule(c *gin.Context) {
	var dataNeeded struct {
		Id string `json:"medication"`
	}
	if err := c.ShouldBindJSON(&dataNeeded); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cronI.Remove(entries[dataNeeded.Id])
	c.IndentedJSON(http.StatusOK, "done")
}
