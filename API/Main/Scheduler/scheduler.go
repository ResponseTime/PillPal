package Scheduler

import (
	"api/Models"
	"api/Utility"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func EditSchedule(c *gin.Context) {}

func SetSchedule(c *gin.Context) {
	var Data Models.Data
	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	scheduledTime, err := time.Parse(time.RFC3339, Data.Time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	var id cron.EntryID
	switch strings.ToLower(Data.Frequency) {
	case "once":
		{
			id, _ = Utility.CronI.AddFunc(fmt.Sprintf("%d %d %d %d *", scheduledTime.Minute(), scheduledTime.Hour(), scheduledTime.Day(), int(scheduledTime.Month())), func() { fmt.Println("reminder for ", Data) })
		}
	case "daily":
		{
			id, _ = Utility.CronI.AddFunc(fmt.Sprintf("%d %d * * *", scheduledTime.Minute(), scheduledTime.Hour()), func() { fmt.Println("reminder for ", Data) })
		}
	case "weekly":
		{
			id, _ = Utility.CronI.AddFunc(fmt.Sprintf("%d %d * * %d", scheduledTime.Minute(), scheduledTime.Hour(), scheduledTime.Weekday()), func() { fmt.Println("reminder for ", Data) })
		}
	}

	Utility.CronI.Start()
	Utility.Entries[Data.Medication+strconv.Itoa(int(id))] = id
	fmt.Println(Utility.CronI.Entries())
	c.IndentedJSON(http.StatusOK, gin.H{"id": Data.Medication + strconv.Itoa(int(id))})
}

func DelSchedule(c *gin.Context) {
	k := c.Params.ByName("id")
	Utility.CronI.Remove(Utility.Entries[k])
	fmt.Println(Utility.CronI.Entries())
	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
