package main

import (
  "github.com/gin-gonic/gin"
  "go_forms/models"
  "go_forms/api"
  "go_forms/src"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/**")
  models.ConnectDatabase()

  r.GET("/", app.HomePage)
  r.GET("/:id", app.FormPage)
  r.GET("/create", app.CreateFormPage)
  r.GET("/add-question", app.Add_Question)
  r.GET("/add-option", app.Add_Option)

  form_routes := r.Group("/forms")
  form_routes.GET("/", api.GetForms)
  form_routes.GET("/get/:title", api.IndexForms)
  form_routes.POST("/create", api.CreateForms)
  form_routes.PATCH("/update/:id", api.UpdateForms)
  form_routes.DELETE("/delete/:id", api.DeleteForms)

  question_routes := r.Group("/questions")
  question_routes.GET("/", api.GetQuestions)
  question_routes.GET("/get/:question", api.IndexQuestions)
  question_routes.POST("/create", api.CreateQuestions)
  question_routes.PATCH("/update/:id", api.UpdateQuestions)
  question_routes.DELETE("/delete/:id", api.DeleteQuestions)

  option_routes := r.Group("/options")
  option_routes.GET("/", api.GetOptions)
  option_routes.GET("/get/:option_text", api.IndexOptions)
  option_routes.POST("/create", api.CreateOptions)
  option_routes.PATCH("/update/:id", api.UpdateOptions)
  option_routes.DELETE("/delete/:id", api.DeleteOptions)
  r.Run()
}
