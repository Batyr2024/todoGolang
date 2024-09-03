package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model // add ID, created_at etc.
	text string`json`

}