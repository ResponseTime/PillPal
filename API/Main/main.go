package main

import (
	"fmt"
	"net/http"

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
	router.POST("api/set/:item", setSchedule)
	router.POST("api/del/:item", delSchedule)
	router.POST("api/test", testRoute)
	router.Run("localhost:8080")

}

func setSchedule(c *gin.Context) {
	item := c.Param("item")
	id, _ := cronI.AddFunc("@every 1m", func() { fmt.Println("reminder for ", item) })
	entries[item] = id
	cronI.Start()
	st := struct {
		Item string
		Set  bool
	}{Item: item, Set: true}
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
