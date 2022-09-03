package app

import (
	"go_forms/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
  forms := api.Fetch_Forms()
  c.HTML(http.StatusOK, "forms_list.html", gin.H{"data": forms})
}
func FormPage(c *gin.Context) {
  form := api.Fetch_Forms_Index(c.Param("id"))
  questions := api.Fetch_Question_Forms(int(form.ID))
  c.HTML(http.StatusOK, "forms_index.html", gin.H{"form": form, "questions": questions})
}

func CreateFormPage(c *gin.Context) {
  c.HTML(http.StatusOK, "form_create.html", gin.H{})
}
