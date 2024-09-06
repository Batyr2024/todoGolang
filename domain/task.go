package domain

type Task struct {
	ID uint   `json:"id" gorm:"unique;not null"`
	Text string`json:"text"`
	Checked bool`json:"checked"`
}