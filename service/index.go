package service

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// IndexHandler godoc
// @Summary ping example
// @Schemes
// @Description index api page
// @Tags index
// @Accept json
// @Produce json
// @Success 200 {string} Hello
// @Router /index [get]
func IndexHandler(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Hello"})
}
