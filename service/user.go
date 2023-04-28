package service

import (
	"github.com/gin-gonic/gin"
	"im/models"
)

func GetAllUsers(context *gin.Context) {
	data := make([]*models.User, 10)
	data = models.GetUsers()
	context.JSON(200, gin.H{"message": data})
}
