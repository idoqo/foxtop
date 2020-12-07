package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/idoko/foxtop/db"
	"gitlab.com/idoko/foxtop/mozurl"
	"log"
	"net/http"
)

var database db.Database

func InitRouter(db db.Database) *gin.Engine {

	database = db

	r := gin.Default()
	r.GET("/hosts", func(c *gin.Context) {
		hosts, err := database.AllHosts()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, hosts)
		}
	})

	r.GET("/hosts/:host", func(c *gin.Context) {
		hostname := c.Param("host")
		host := &mozurl.MozHost{Host: hostname}
		err := db.URLsForHost(host)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"host": host})
		}
	})

	r.GET("/protocols", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "not implemented"})
	})
	r.GET("/timekeeper", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "not implemented"})
	})
	r.GET("/bookmarks", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "not implemented"})
	})
	return r
}
