package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model // add ID, created_at etc.
	Text string`json:"text"`
	Checked

}