package app

import (
	"go_forms/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateForm(c *gin.Context) {
  form := api.Fetch_Forms_Index(c.Param("id"))
  questions := api.Fetch_Question_Forms(int(form.ID))
  c.HTML(http.StatusOK, "forms_index.html", gin.H{"form": form, "questions": questions})
}

func Add_Question(c *gin.Context) {
  c.HTML(http.StatusOK, "question_form.html", gin.H{})
}
func Add_Option(c *gin.Context) {
  c.HTML(http.StatusOK, "option_form.html", gin.H{})
}

