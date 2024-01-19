package Router

import (
	"api/Scheduler"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/api/main")
	router.POST("/api/set", Scheduler.SetSchedule)
	router.DELETE("/api/del/:id", Scheduler.DelSchedule)
	router.PATCH("/api/edit", Scheduler.EditSchedule)
	return router
}
