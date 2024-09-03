package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model // add: ID, CreatedAt, UpdatedAt, DeletedAt
	Text string`json:"text"`
	isChecked bool`json:"isCheck"`
}