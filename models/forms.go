package models

type Forms struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title string `json:"title"`
  Description string `json:"description"`
  Question []Question `json:"question"`
}

type Question struct {
  ID uint `json:"id" gorm:"primary_key"`
  Question string `json:"question_name"`
  Options *Options `json:"options"`
}

type Options struct {
  ID uint `json:"id" gorm:"primary_key"`
  Option_text string `json:"option"`
  Is_correct bool `json:"is_correct"`
}
 
type Create_Form_Input struct {
  Title string `json:"title"`
  Description string `json:"description"`
  QuestionID string `json:"question_id"`
}
