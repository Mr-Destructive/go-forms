package api

import (
	"go_forms/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fetch_Options() []models.Options{
  var Options []models.Options
  models.DB.Find(&Options)
  return Options
}

func Fetch_Options_Index(option_text string) models.Options{
  var Options models.Options

  err := models.DB.Where("Option_text= ?", option_text).Find(&Options).Error
      if err != nil {
        panic(err)
      }
  return Options
}

func GetOptions(c *gin.Context) {
  Options := Fetch_Options()
  c.JSON(http.StatusOK, gin.H{"data": Options})
}

func CreateOptions(c *gin.Context) {
  var input models.Create_Option_Input

  if input.Option_text == ""{
    if c.PostForm("is_correct") == "true"{
      input.Is_correct = true
    } else if c.PostForm("is_correct") == "false" {
      input.Is_correct = false
    }
    input.Option_text = c.PostForm("option")
  }
  //if err := c.ShouldBindJSON(&input); err != nil {
  //  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  //  return
  //}

  Options := models.Options{Option_text: input.Option_text}
  models.DB.Create(&Options)

  c.JSON(http.StatusOK, gin.H{"data": Options})
}

func IndexOptions(c *gin.Context) {
 var Options models.Options

  if err := models.DB.Where("Option_text= ?", c.Param("option_text")).First(&Options).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": Options})
}

func UpdateOptions(c *gin.Context) {
  var Options models.Options
  if err := models.DB.Where("id = ?", c.Param("id")).First(&Options).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  var input models.Create_Option_Input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&Options).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": Options})
}

func DeleteOptions(c *gin.Context) {
  var Options models.Options
  if err := models.DB.Where("id = ?", c.Param("id")).First(&Options).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&Options)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
