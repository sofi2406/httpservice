package main

import (
	"assignment/httpservice/database"

	"github.com/gin-gonic/gin"
)

func GetHttpService() *gin.Engine {
	r := gin.Default()

	r.GET("/data/:id", func(c *gin.Context) {
		data, err := database.GetDataSet(c.Param("id"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, data)
	})
	return r
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/data/:id", func(c *gin.Context) {
		data, err := database.GetDataSet(c.Param("id"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, data)
	})
	r.Run(":3000")
}
