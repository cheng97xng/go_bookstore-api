package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	now := time.Now()
	n := now.Format("2006-1-02T15:04:05.000Z")
	fmt.Println("now date:", n)
	c.String(http.StatusOK, "Cheng..")
}
