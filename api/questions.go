package api

import (
	"go_forms/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Fetch_Questions() []models.Question {
	var Questions []models.Question
	models.DB.Find(&Questions)
	return Questions
}

func Fetch_Question_Forms(id int) []models.Question {
	var Questions []models.Question

	err := models.DB.Where("form_id= ?", id).Find(&Questions).Error
	if err != nil {
		panic(err)
	}
	return Questions
}

func Fetch_Questions_Index(question string) models.Question {
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
	if input.OptionID != "" || input.Question == "" {
		input.Question = c.PostForm("question")
		input.OptionID = c.PostForm("option_ids")
	}
	//if err := c.ShouldBindJSON(&input); err != nil {
	//  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//  return
	//}

	var option_list []models.Options
	for _, option := range strings.Split(input.OptionID, ",") {
		if o, err := strconv.Atoi(option); err == nil {
			var option_ref models.Options
			err := models.DB.Where("id= ?", o).First(&option_ref).Error
			if err != nil {
				log.Panic(err)
			}
			option_list = append(option_list, option_ref)
			//question_ref.Option...)
		}
	}
	Questions := models.Question{Question: input.Question, Option: option_list}
	models.DB.Create(&Questions)

	c.JSON(http.StatusOK, gin.H{"data": Questions.ID})
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

	var option_list []models.Options
	for _, option := range strings.Split(input.OptionID, ",") {
		if o, err := strconv.Atoi(option); err == nil {
			var option_ref models.Options
			err := models.DB.Where("id= ?", o).First(&option_ref).Error
			if err != nil {
				log.Panic(err)
			}
			option_list = append(option_list, option_ref)
			//question_ref.Option...)
		}
	}

	models.DB.Model(&Questions).Updates(models.Question{Question: input.Question, Option: option_list})

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
