package models

type Blog struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Text    string `json:"blog_text"`
	User_ID uint   `json:"user_id"`
}
