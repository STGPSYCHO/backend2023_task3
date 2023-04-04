package models

type Blog struct {
	ID      uint      `json:"id" gorm:"primary_key"`
	Title   string    `json:"blog_title"`
	Text    string    `json:"blog_text"`
	Comment []Comment `gorm:"ForeignKey:BlogID"`
}
