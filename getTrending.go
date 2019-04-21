package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTrending(c *gin.Context) {

	date := c.Param("date")

	db := connectDB()
	defer db.Close()

	list := findByDate(date, db)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"list": list,
	})

}