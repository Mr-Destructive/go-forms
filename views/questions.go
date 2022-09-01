package api

import (
	"go_forms/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Fetch_Questions() []models.Question{
  var Questions []models.Question
  models.DB.Find(&Questions)
  return Questions
}

func Fetch_Questions_Index(question string) models.Question{
  var Questions models.Question

  err := models.DB.Where("Question= ?", question).Find(&Questions).Error
      if err != nil {
        panic(err)
      }
  return Questions
}

func GetQuestions(c *gin.Context) {
  Questions := Fetch_Questions()
  c.JSON(http.StatusOK, gin.H{"data": Questions})
}

func CreateQuestions(c *gin.Context) {
  var input models.Create_Question_Input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var option_list []models.Options
  var option_ref models.Options
  var question_ref models.Question
  log.Print(strings.Split(input.OptionID, ","))
  for _, option := range strings.Split(input.OptionID, ","){
    log.Print(question_ref.Option)
    if o, err := strconv.Atoi(option); err == nil {
      log.Print(models.DB.Select("question_id").Where("id= ?", o).First(&models.Options{}))
      err := models.DB.Where("id= ?", o).Preload("Option").First(&question_ref).Error
      log.Print(option_ref.Question_ID)
      if err != nil {
        log.Panic(err)
      }
      option_list = append(option_list, question_ref.Option...)
      log.Print(option_list)
     }
  }
  log.Print(option_list)
  Questions := models.Question{Question: input.Question, Option: option_list}
  models.DB.Create(&Questions)

  c.JSON(http.StatusOK, gin.H{"data": Questions})
}

func IndexQuestions(c *gin.Context) {
 var Questions models.Question

  if err := models.DB.Where("Question= ?", c.Param("question")).First(&Questions).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": Questions})
}

func UpdateQuestions(c *gin.Context) {
  var Questions models.Question
  if err := models.DB.Where("id= ?", c.Param("id")).First(&Questions).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  var input models.Create_Question_Input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&Questions).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": Questions})
}

func DeleteQuestions(c *gin.Context) {
  var Questions models.Question
  if err := models.DB.Where("id= ?", c.Param("id")).First(&Questions).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&Questions)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
