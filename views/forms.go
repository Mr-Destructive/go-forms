package api

import(
	"go_forms/models"
  "net/http"
	"github.com/gin-gonic/gin"
)

func Fetch_Forms() []models.Forms{
  var Forms []models.Forms
  models.DB.Find(&Forms)
  return Forms
}

func Fetch_Forms_Index(title string) models.Forms{
  var Forms models.Forms

  err := models.DB.Where("Title= ?", title).Find(&Forms).Error
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

  Forms := models.Forms{Title: input.Title, Description: input.Description}
  models.DB.Create(&Forms)

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

  models.DB.Model(&Forms).Updates(input)

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
