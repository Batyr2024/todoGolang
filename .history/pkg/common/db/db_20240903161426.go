package db

import (
	"log"
	"net/url"

	"github.com/Batyr2024/todoGolang/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string)