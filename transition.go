package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func transition(c *gin.Context) {
	c.HTML(http.StatusOK, "dateBy.html", gin.H{})
}