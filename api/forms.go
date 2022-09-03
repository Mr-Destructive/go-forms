package api

import(
	"go_forms/models"
  "net/http"
  "strconv"
  "strings"
  "log"
	"github.com/gin-gonic/gin"
)

func Fetch_Forms() []models.Forms{
  var Forms []models.Forms
  models.DB.Find(&Forms)
  return Forms
}

func Fetch_Forms_Index(id string) models.Forms{
  var Forms models.Forms

  err := models.DB.Where("id= ?", id).Find(&Forms).Error
      if err != nil {
        panic(err)
      }
  return Forms
}

func GetForms(c *gin.Context) {
  Forms := Fetch_Forms()
  c.JSON(http.StatusOK, gin.H{"data": Forms})
}

func CreateForms(c *gin.Context) {
  var input models.Create_Form_Input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }


  log.Print(input)
  var question_list []models.Question
  for _, question := range strings.Split(input.QuestionID, ","){
    if o, err := strconv.Atoi(question); err == nil {
      var question_ref models.Question
      err := models.DB.Where("id= ?", o).First(&question_ref).Error
      if err != nil {
        log.Panic(err)
      }
      question_list = append(question_list, question_ref)
     }
  }
  Forms := models.Forms{Title: input.Title, Description: input.Description, Question: question_list}
  models.DB.Create(&Forms)
  for _, q := range question_list{
    var Questions models.Question
    if err := models.DB.Where("id= ?", q.ID).First(&Questions).Error; err != nil{
      return 
    }
    models.DB.Model(&Questions).Updates(models.Question{Form_ID: Forms.ID})
  }

  c.JSON(http.StatusOK, gin.H{"data": Forms})
}

func IndexForms(c *gin.Context) {
 var Forms models.Forms
  if err := models.DB.Where("Title= ?", c.Param("title")).First(&Forms).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": Forms})
}

func UpdateForms(c *gin.Context) {
  var Forms models.Forms
  if err := models.DB.Where("id = ?", c.Param("id")).First(&Forms).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
  var input models.Create_Form_Input
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var question_list []models.Question
  for _, question := range strings.Split(input.QuestionID, ","){
    if o, err := strconv.Atoi(question); err == nil {
      var question_ref models.Question
      var Questions models.Question
      if err := models.DB.Where("id= ?", o).First(&Questions).Error; err != nil{
        return 
      }
      models.DB.Model(&Questions).Updates(models.Question{Form_ID: Forms.ID})
      err := models.DB.Where("id= ?", o).First(&question_ref).Error
      if err != nil {
        log.Panic(err)
      }
      question_list = append(question_list, question_ref)
     }
  }

  models.DB.Model(&Forms).Updates(models.Forms{Title: input.Title, Description: input.Description, Question: question_list})

  c.JSON(http.StatusOK, gin.H{"data": Forms})
}

func DeleteForms(c *gin.Context) {
  var Forms models.Forms
  if err := models.DB.Where("id = ?", c.Param("id")).First(&Forms).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&Forms)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
