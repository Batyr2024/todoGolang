package tasks

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Batyr2024/todoGolang/pkg/common/models"
)

type AddTaskRequestBody struct{
	Text string `json:"text"`
}

