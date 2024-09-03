package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model{ID } // add: ID, CreatedAt, UpdatedAt, DeletedAt
	Text string`json:"text"`
	Checked bool`json:"check"`
}