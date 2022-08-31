package main

import (
  "github.com/gin-gonic/gin"
  "go_forms/models"
  "go_forms/views"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/**")
  models.ConnectDatabase()
  r.GET("/", api.Hello_World)
  r.GET("/hello", api.Render_Hello)
  r.POST("/create", api.CreateForms)
  r.GET("/forms", api.GetForms)
  r.GET("/forms/:title", api.IndexForms)
  r.GET("/update/:title", api.UpdateForms)
  r.GET("/delete/:title", api.DeleteForms)
  r.Run()
}
