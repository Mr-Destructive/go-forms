package api

import (
  "net/http"
	"github.com/gin-gonic/gin"
)
func Hello_World(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"data":"hello world"})
}
func Render_Hello(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{"data": "hello world",})
}

