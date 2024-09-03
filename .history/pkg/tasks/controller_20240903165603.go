package tasks

import(
	"gorm.io/gorm"
)

type handler struct{
	DB *gorm.DB
}

func RegisterRoutes