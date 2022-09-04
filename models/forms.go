package models

import "github.com/jinzhu/gorm"

type Forms struct {
	gorm.Model
	//ID     uint   `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Question    []Question `json:"question_list"`
}

type Question struct {
	gorm.Model
	//ID uint `json:"id" gorm:"primary_key"`
	Question string    `json:"question_name"`
	Form_ID  uint      `json:"form_id"`
	Option   []Options `json:"option_list"`
}

type Options struct {
	gorm.Model
	//ID uint `json:"id" gorm:"primary_key"`
	Option_text string `json:"option"`
	Is_correct  bool   `json:"is_correct"`
	Question_ID uint   `json:"question_id"`
}

type Create_Form_Input struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuestionID  string `json:"question_ids"`
}

type Create_Question_Input struct {
	Question string `json:"question"`
	OptionID string `json:"option_ids"`
}

type Create_Option_Input struct {
	Option_text string `json:"option_text"`
	Is_correct  bool   `json:"is_correct"`
}
