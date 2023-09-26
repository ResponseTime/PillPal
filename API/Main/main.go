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
	router.GET("api/set/:item", setSchedule)
	router.GET("api/del/:item", delSchedule)
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
